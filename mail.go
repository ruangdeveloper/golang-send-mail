package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

const (
	smtpHost    = "sandbox.smtp.mailtrap.io"
	smtpPort    = 587
	smtpUser    = "43374cf52e7c16"
	smtpPass    = "5c9a4af620d09f"
	mailFrom    = "no-reply@example.com"
	mailReplyTo = "hello@example.com"
)

func InitDialer() *gomail.Dialer {
	dialer := gomail.NewDialer(
		smtpHost,
		smtpPort,
		smtpUser,
		smtpPass,
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer
}

func SendMail(
	to string,
	subject string,
	message string,
) error {
	dialer := InitDialer()

	mail := gomail.NewMessage()

	mail.SetHeader("From", mailFrom)
	mail.SetHeader("To", to)
	mail.SetHeader("Reply-To", mailReplyTo)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", message)

	return dialer.DialAndSend(mail)
}
