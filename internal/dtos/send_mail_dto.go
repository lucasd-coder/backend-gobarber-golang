package dtos

type SendMailDTO struct {
	From    string   `form:"from" default:"Equipe GoBarber"`
	To      []string `form:"to"`
	Message []byte   `form:"message" binding:"requird"`
}
