package backends



type NotificationPayload struct {
	To   string `json:"to"`
	Data struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"data"`
}

func NewNotificationPayload(to string, data map[string]string) *NotificationPayload{
	payload := &NotificationPayload{
		To: to,
	}
	payload.Data.Body = data["body"]
	payload.Data.Title = data["title"]
	return payload
}


type PushNotificationSender interface {
	SendNotification(NotificationPayload) error
}
