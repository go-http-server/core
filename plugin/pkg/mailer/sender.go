package mailer

import (
	"fmt"
	"html/template"

	"github.com/google/uuid"
	"github.com/wneessen/go-mail"
)

type UserReceive struct {
	Username     string
	EmailAddress string
	Code         uuid.UUID
	Fullname     string
}

type EmailSender interface {
	SendWithTemplate(subject, pathTemplate string, to UserReceive, attachFiles []string) error
}

type GmailSender struct {
	name             string
	emailAddress     string
	emailAppPassword string
}

func NewGmailSender(name, emailAddress, emailAppPassword string) EmailSender {
	return &GmailSender{
		name:             name,
		emailAddress:     emailAddress,
		emailAppPassword: emailAppPassword,
	}
}

func (sender GmailSender) SendWithTemplate(subject, pathTemplate string, to UserReceive, attachFiles []string) error {
	htmlTemplate, err := template.ParseFiles(pathTemplate)
	if err != nil {
		return fmt.Errorf("Error with path template file: %s", err)
	}
	message := mail.NewMsg()

	if err := message.EnvelopeFrom(sender.emailAddress); err != nil {
		return fmt.Errorf("failed to set ENVELOPE FROM address: %s", err)
	}
	if err := message.FromFormat(sender.name, sender.emailAddress); err != nil {
		return fmt.Errorf("failed to set formatted FROM address: %s", err)
	}
	if err := message.AddToFormat(to.Fullname, to.EmailAddress); err != nil {
		return fmt.Errorf("failed to set formatted TO address: %s", err)
	}

	message.SetMessageID()
	message.SetDate()

	message.Subject(subject)
	if err := message.AddAlternativeHTMLTemplate(htmlTemplate, to); err != nil {
		return fmt.Errorf("failed to add HTML template to mail body: %s", err)
	}

	for _, filePath := range attachFiles {
		message.AttachFile(filePath, mail.WithFileDescription("From Go with love"))
	}

	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithUsername(sender.emailAddress),
		mail.WithPassword(sender.emailAppPassword),
	)
	if err != nil {
		return fmt.Errorf("Cannot create client mailer: %s", err)
	}

	if err = client.DialAndSend(message); err != nil {
		return fmt.Errorf("Cannot delivery mail: %s", err)
	}
	return nil
}
