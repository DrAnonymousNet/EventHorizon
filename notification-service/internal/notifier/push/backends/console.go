package backends

import (
	"fmt"
	"sync"
)

type ConsolePushNotifier struct {
	lock sync.Mutex
}

func NewConsolePushNotificationSender() *ConsolePushNotifier {
	return &ConsolePushNotifier{}
}

func (cpn *ConsolePushNotifier) SendNotification(payload NotificationPayload) error {
	cpn.lock.Lock()
	defer cpn.lock.Unlock()
	fmt.Printf("Notification payload: %v", payload)
	return nil
}
