package services

import (
	"log"
	"net/smtp"
)

type MailOrigin struct {
	from     string
	password string
}

////////////send mail
func SendMail(body string, recipient string, origin MailOrigin) (er bool) {

	msg := "From: " + origin.from + "\n" +
		"To: " + recipient + "\n" +
		"Subject: Your verification code is : \n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", origin.from, origin.password, "smtp.gmail.com"),
		origin.from, []string{recipient}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	log.Print("verification code sent to : " + recipient)
	return true
}

////////////////////////////////////////////////////////
