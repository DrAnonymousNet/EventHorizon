package email_backend

import (
	"fmt"
	"sync"
)

//Interface for email backend implementations.

// ConsoleEmailSender is an implementation of EmailSender that writes emails to the console
type ConsoleEmailSender struct {
	lock sync.Mutex
}

// NewConsoleEmailSender creates a new instance of ConsoleEmailSender
func NewConsoleEmailSender() *ConsoleEmailSender {
	return &ConsoleEmailSender{}
}

// SendEmail writes the email details to the console in a thread-safe way
func (ces *ConsoleEmailSender) SendEmail(email Email) error {
	ces.lock.Lock()
	defer ces.lock.Unlock()

	// Simulate writing the email message to the console
	fmt.Printf("From: %s\n", email.From)
	fmt.Printf("To: %s\n", email.To)
	fmt.Printf("Subject: %s\n", email.Subject)
	fmt.Printf("Body:\n%s\n", email.Body)
	fmt.Println("------------------------------------------------------------")

	return nil
}

func (ces *ConsoleEmailSender) SendEmailWithAttachment(email Email) error {
	ces.lock.Lock()
	defer ces.lock.Unlock()

	fmt.Printf("From: %s\n", email.From)
	fmt.Printf("To: %s\n", email.To)
	fmt.Printf("Subject: %s\n", email.Subject)
	fmt.Printf("Body:\n%s\n", email.Body)
	fmt.Printf("Attachment: %s\n", email.Attachments[0]["filename"])
	fmt.Println("------------------------------------------------------------")

	return nil
}