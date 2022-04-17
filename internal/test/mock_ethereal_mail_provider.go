package test

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model/external"

	"github.com/stretchr/testify/mock"
)

type MockEtherealMailProvider struct {
	mock.Mock
}

func (mock *MockEtherealMailProvider) SendMail(authSmtp *external.AuthSmtpSendEmail, dto *dtos.SendMailDTO) error {
	return nil
}
