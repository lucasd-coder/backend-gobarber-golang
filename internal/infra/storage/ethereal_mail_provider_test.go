package storage_test

import (
	"net/smtp"
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/storage"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model/external"
	"github.com/stretchr/testify/assert"
)

func TestEtherealMailSendMailSuccessfully(t *testing.T) {
	authSmtp := &external.AuthSmtpSendEmail{
		Host:     "me@example.com",
		Port:     "8000",
		Password: "12345",
		Username: "test",
	}

	body := "Hello World"
	dto := &dtos.SendMailDTO{
		From:    "me@example.com",
		To:      []string{"me@example.com"},
		Message: []byte(body),
	}

	f := mockSend(nil)

	testEtherealMail := storage.EtherealMailProvider{Send: f}

	err := testEtherealMail.SendMail(authSmtp, dto)

	assert.Nil(t, err)
	assert.NoError(t, err)
}

func mockSend(errToReturn error) func(string, smtp.Auth, string, []string, []byte) error {
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return errToReturn
	}
}
