package integration

import (
	"api/config"
	"api/internal/constants"
	"api/internal/controllers/http"
	"api/internal/infrastructure/cryptoProviders"
	"api/internal/infrastructure/repository/fileStorage"
	"api/internal/usecases"
	"api/internal/usecases/details"
	"api/internal/usecases/usecases_contracts"
	mock_usecases_contracts "api/internal/usecases/usecases_contracts/mocks"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

const TestDataPath = "../data/data.json"

type IntegrationTestSuite struct {
	suite.Suite

	cfg         *config.Config
	cryptoChain details.CryptoChain
	handler     *http.Handler
	useCases    *usecases.UseCases
	repos       *usecases_contracts.Repository

	mailerMock *mock_usecases_contracts.MockMailer
}

func TestIntegrationTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic(err)
	}

	dataToWrite := []byte(`{"emails":[]}`)
	err = os.WriteFile(TestDataPath, dataToWrite, constants.WriteFilePerm)
	if err != nil {
		s.FailNowf("unable to setup data fileStorage", err.Error())
	}

	s.initDeps()
}

func (s *IntegrationTestSuite) TearDownSuite() {
	err := os.Truncate(TestDataPath, 0)
	if err != nil {
		s.FailNowf("unable to clear data fileStorage", err.Error())
	}
}

func (s *IntegrationTestSuite) initDeps() {
	mockController := gomock.NewController(s.T())
	s.mailerMock = mock_usecases_contracts.NewMockMailer(mockController)

	s.cfg = config.GetConfig()
	coinMarketCapProviderCreator := crypto_providers.NewCoinMarketCapProviderCreator(s.cfg)
	binanceProviderCreator := crypto_providers.NewBinanceProviderCreator(s.cfg)
	coinAPIProviderCreator := crypto_providers.NewCoinAPIProviderCreator(s.cfg)
	coinbaseProviderCreator := crypto_providers.NewCoinbaseProviderCreator(s.cfg)

	coinMarketCapProvider := coinMarketCapProviderCreator.CreateCryptoProvider()
	binanceProvider := binanceProviderCreator.CreateCryptoProvider()
	coinAPIProvider := coinAPIProviderCreator.CreateCryptoProvider()
	coinbaseProvider := coinbaseProviderCreator.CreateCryptoProvider()

	coinMarketCapChain := details.NewBaseCryptoChain(coinMarketCapProvider)
	binanceChain := details.NewBaseCryptoChain(binanceProvider)
	coinAPIChain := details.NewBaseCryptoChain(coinAPIProvider)
	coinbaseChain := details.NewBaseCryptoChain(coinbaseProvider)

	coinMarketCapChain.SetNext(binanceChain)
	binanceChain.SetNext(coinAPIChain)
	coinAPIChain.SetNext(coinbaseChain)

	s.cryptoChain = coinMarketCapChain

	s.repos = initRepos(TestDataPath)
	s.useCases = initUseCases(s.repos, s.cryptoChain, s.mailerMock, s.cfg)
	s.handler = http.NewHandler(s.useCases)
}

func initRepos(filePath string) *usecases_contracts.Repository {
	emailSub := fileStorage.NewEmailSubscriptionRepository(filePath)

	return fileStorage.NewRepository(emailSub)
}

func initUseCases(repositories *usecases_contracts.Repository, cryptoChain details.CryptoChain, mailer usecases_contracts.Mailer, cfg *config.Config) *usecases.UseCases {
	getRate := details.NewCachedRateGetter(usecases.NewGetRateUseCase(cryptoChain), cfg.Cache.RateCacheTTL)
	sendEmails := usecases.NewSendEmailsUseCase(repositories.EmailSubscriptionRepo, mailer, getRate)
	subscribeEmails := usecases.NewSubscribeEmailUseCase(repositories.EmailSubscriptionRepo)

	return usecases.NewUseCases(getRate, sendEmails, subscribeEmails)
}
