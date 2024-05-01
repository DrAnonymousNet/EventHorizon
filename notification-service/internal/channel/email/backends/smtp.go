package email_backend

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"strings"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)

type SMTPEmailSender struct {
	Server config.SMTPServer
}

func NewSMTPEmailSender() *SMTPEmailSender {
	return &SMTPEmailSender{
		Server: *config.SmtpSetting,
	}
}

func (s *SMTPEmailSender) SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", s.Server.From, s.Server.PassWord, s.Server.Host)
	err := smtp.SendMail(s.Server.Host+":"+s.Server.Port, auth, s.Server.From, email.To, []byte(email.Body))
	if err != nil {
		log.Println("Error sending email: ", err)
		return err
	}
	return nil
}

// SendEmailWithAttachment sends an email with an attachment
func (s *SMTPEmailSender) SendEmailWithAttachment(email Email) error {
	// Connect to the SMTP server
	auth := smtp.PlainAuth("", s.Server.From, s.Server.PassWord, s.Server.Host)
	addr := fmt.Sprintf("%s:%s", s.Server.Host, s.Server.Port)

	// Create a new multipart writer for the email body
	var emailBody bytes.Buffer
	writer := multipart.NewWriter(&emailBody)

	// Add CC and BCC recipients to the list of recipients for smtp.SendMail
    // Note: BCC recipients are not added to the email headers to keep them hidden
    recipients := append(email.To, email.Cc...)
    recipients = append(recipients, email.Bcc...)

	// Write the headers
	headers := map[string]string{
		"From":         email.From,
		"To":           strings.Join(email.To, ", "),
		"Cc":           strings.Join(email.Cc, ", "),
		"Subject":      email.Subject,
		"MIME-Version": "1.0",
		"Content-Type": fmt.Sprintf("multipart/mixed; boundary=%s", writer.Boundary()),
	}
	for k, v := range headers {
		emailBody.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	emailBody.WriteString("\r\n")

	// Write the email body
	var bodyContentType string
	if email.HTMLBody != "" {
		bodyContentType = "text/html"
		email.Body = email.HTMLBody
	} else {
		bodyContentType = "text/plain"
	}
	part, err := writer.CreatePart(textproto.MIMEHeader(map[string][]string{"Content-Type": {bodyContentType}}))
	if err != nil {
		return err
	}

	_, err = part.Write([]byte(email.Body))
	if err != nil {
		return err
	}

	for _, attachment := range email.Attachments {
		// Attach the file
		attachmentBytes := attachment["content"]
		if err != nil {
			return err
		}
		attachmentPart, err := writer.CreatePart(map[string][]string{
			"Content-Type":              {"application/octet-stream"},
			"Content-Disposition":       {fmt.Sprintf("attachment; filename=\"%s\"", string(attachment["filename"]))},
			"Content-Transfer-Encoding": {"base64"},
		})
		if err != nil {
			return err
		}

		base64Encoder := base64.NewEncoder(base64.StdEncoding, attachmentPart)
		base64Encoder.Write(attachmentBytes)
		base64Encoder.Close()
	}


	// Close the multipart writer
	writer.Close()

	// Send the email
	return smtp.SendMail(addr, auth, email.From, recipients, emailBody.Bytes())
}
