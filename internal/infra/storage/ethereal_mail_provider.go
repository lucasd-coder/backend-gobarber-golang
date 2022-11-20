package storage

import (
	"net/smtp"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model/external"
)

type EtherealMailProvider struct {
	Send func(string, smtp.Auth, string, []string, []byte) error
}

func NewEtherealMailProvider() *EtherealMailProvider {
	return &EtherealMailProvider{Send: smtp.SendMail}
}

func (etherealMail *EtherealMailProvider) SendMail(authSmtp *external.AuthSmtpSendEmail, dto *dtos.SendMailDTO) error {
	auth := smtp.PlainAuth(dto.From, authSmtp.Username, authSmtp.Password, authSmtp.Host)

	err := etherealMail.Send(authSmtp.Host+":"+authSmtp.Port, auth, dto.From, dto.To, dto.Message)

	return err
}
