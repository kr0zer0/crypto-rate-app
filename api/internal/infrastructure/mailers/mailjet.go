package mailers

import (
	"api/config"
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go"
)

type MailjetMailer struct {
	cfg           *config.Config
	mailjetClient *mailjet.Client
}

func NewMailjetMailer(cfg *config.Config, mailjetClient *mailjet.Client) *MailjetMailer {
	return &MailjetMailer{
		cfg:           cfg,
		mailjetClient: mailjetClient,
	}
}

func (r *MailjetMailer) SendToList(emails []string, message string) error {
	if len(emails) == 0 {
		return nil
	}

	sendingList := r.FormSendingList(emails, message)

	email := &mailjet.InfoSendMail{
		Messages: sendingList,
	}

	_, err := r.mailjetClient.SendMail(email)
	if err != nil {
		return err
	}

	return nil
}

func (r *MailjetMailer) FormSendingList(emails []string, message string) []mailjet.InfoSendMail {
	var sendingList []mailjet.InfoSendMail
	for _, email := range emails {
		info := mailjet.InfoSendMail{
			FromEmail: r.cfg.EmailSending.SenderAddress,
			FromName:  "BTC app",
			Recipients: []mailjet.Recipient{
				{
					Email: email,
				},
			},
			Subject:  "BTC exchange rate",
			HTMLPart: fmt.Sprintf("<h3>%s</h3>", message),
		}
		sendingList = append(sendingList, info)
	}

	return sendingList
}
