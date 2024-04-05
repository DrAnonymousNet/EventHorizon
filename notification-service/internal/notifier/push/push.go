package notifier

import (
	"github.com/dranonymousnet/eventhorizon/internal/config"
	"github.com/dranonymousnet/eventhorizon/internal/notifier/push/backends"
)

func GetPushNotificationSender() backends.PushNotificationSender {
	if config.AppSetting.PushNotificationBackend == "fcm" || config.AppSetting.Environment == "production" {
		return backends.NewFCMNotificationSender()
	} else {
		return backends.NewConsolePushNotificationSender()
	}
}

func SendPushNotification(payload backends.NotificationPayload) error {
	sender := GetPushNotificationSender()
	return sender.SendNotification(payload)
}
