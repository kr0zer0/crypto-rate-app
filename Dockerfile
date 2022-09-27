FROM golang:latest

WORKDIR /app

COPY api/go.mod .
COPY api/go.sum .

RUN go mod download

COPY api/ .

RUN cd main/; go build

EXPOSE 8000

CMD ["main/main"]