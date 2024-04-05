package email_backend

import (
	"os"
	"path/filepath"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)

type Email struct {
	From        string
	To          []string
	Subject     string
	Body        string
	Cc          []string
	Bcc         []string
	HTMLBody    string
	Attachments []map[string][]byte
}

func (e *Email) AddAttachment(attachmentPath string) error {
	attachmentBytes, err := os.ReadFile(attachmentPath)
	if err != nil {
		return err
	}
	attachment := map[string][]byte{
		"filename":[]byte(filepath.Base(attachmentPath)),
		"content":attachmentBytes,
	}
	e.Attachments = append(e.Attachments, attachment)
	return nil
}

func NewEmail(from string, to []string, subject string, cc []string, bcc []string, body string, attachments []map[string][]byte) *Email {
	if len(from) == 0 {
		from = config.AppSetting.FromEmail
	}
	return &Email{
		From:        from,
		To:          to,
		Subject:     subject,
		Body:        body,
		Cc:          cc,
		Attachments: attachments,
	}
}

type EmailSender interface {
	SendEmail(Email) error
	SendEmailWithAttachment(Email) error
}
