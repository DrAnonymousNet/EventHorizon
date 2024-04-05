package notifier

import (
	"fmt"

	"github.com/dranonymousnet/eventhorizon/internal/config"
	email_backend "github.com/dranonymousnet/eventhorizon/internal/notifier/email/backends"
)

func GetEmailSender() email_backend.EmailSender {
	if config.AppSetting.EmailBackend == "smtp" || config.AppSetting.Environment == "production"{
		return email_backend.NewSMTPEmailSender()
	}else if config.AppSetting.EmailBackend == "console"{
		return email_backend.NewConsoleEmailSender()
	}else {
		fmt.Println("No email backend configured")
		return email_backend.NewSMTPEmailSender()
	}
}

//TODO Add Logging
func SendEmail(email email_backend.Email) error {
	sender := GetEmailSender()
	return sender.SendEmail(email)
}