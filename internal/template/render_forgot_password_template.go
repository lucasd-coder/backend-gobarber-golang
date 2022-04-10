package template

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"text/template"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model/external"
	"backend-gobarber-golang/pkg/logger"
)

type RenderForgotPasswordTemplate struct{}

func NewRenderForgotPasswordTemplate() *RenderForgotPasswordTemplate {
	return &RenderForgotPasswordTemplate{}
}

func (render *RenderForgotPasswordTemplate) Render(variables *external.Variables, email string) *dtos.SendMailDTO {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	var body bytes.Buffer

	path := filepath.Join(basepath, "../template/", filepath.Base("forgot_password.html"))

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	tmpl.Execute(&body, variables)

	msg := buildMessage(body)

	mail := &dtos.SendMailDTO{
		From: "equipe@gobarber.com.br",
		To: []string{
			email,
		},
		Message: msg,
	}

	return mail
}

func buildMessage(body bytes.Buffer) []byte {
	subject := fmt.Sprintf("Subject: %s\r\n", "[GoBarber] Recuperação de senha")
	from := fmt.Sprintf("From: Equipe GoBarber <%s\r\n>", "equipe@gobarber.com.br")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	return []byte(subject + from + mime + body.String())
}
