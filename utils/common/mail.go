package common

import (
	"gopkg.in/mail.v2"
)

func SendMail(from, to, subject, content string) error {
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", content)

	dialer := mail.NewDialer("smtp.qq.com", 465, from, "fguovqmlprkmbijc")
	dialer.StartTLSPolicy = mail.MandatoryStartTLS
	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
