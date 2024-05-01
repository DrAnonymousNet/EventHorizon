package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/dranonymousnet/eventhorizon/internal/store"
)

// NotificationType defines the type of the notification
type NotificationType string

const (
	EmailNotification NotificationType = "email"
	SMSNotification   NotificationType = "sms"
	PushNotification  NotificationType = "push"
)

// NotificationStatus defines the status of the notification
type NotificationStatus string

const (
	Pending NotificationStatus = "pending"
	Sent    NotificationStatus = "sent"
	Failed  NotificationStatus = "failed"
)

type Notification struct {
	gorm.Model
	UserID      uint               // Foreign key for the User
	Type        NotificationType   `gorm:"type:varchar(100)"`                   // The type of the notification (email, SMS, etc.)
	Recipient   string             `gorm:"type:varchar(255)"`                   // The recipient identifier (could be an email, phone number, device ID, etc.)
	Content     string             `gorm:"type:text"`                           // The content of the notification
	Status      NotificationStatus `gorm:"type:varchar(100);default:'pending'"` // The status of the notification
	ScheduledAt *time.Time         `gorm:"index"`                               // Optional: specifies when the notification is scheduled to be sent
	SentAt      *time.Time         `gorm:"index"`                               // When the notification was actually sent
}

func (n *Notification) IsSent() bool {
	return n.Status == Sent
}

func (n *Notification) IsFailed() bool {
	return n.Status == Failed
}

func (n *Notification) IsPending() bool {
	return n.Status == Pending
}

// MarkAsSent interact with the database and marks the notification as sent and sets the sent time
func (n *Notification) MarkAsSent() {
	n.Status = Sent
	now := time.Now()
	n.SentAt = &now
	store.DB.Save(n)
}

// MarkAsFailed interact with the database and marks the notification as failed
func (n *Notification) MarkAsFailed() {
	n.Status = Failed
	store.DB.Save(n)
}

// MarkAsPending interact with the database and marks the notification as pending
func (n *Notification) MarkAsPending() {
	n.Status = Pending
	store.DB.Save(n)
}

// GetNotificationsByStatus retrieves notifications from the database based on the status
func GetNotificationsByStatus(status NotificationStatus) []Notification {
	var notifications []Notification
	store.DB.Where("status = ?", status).Find(&notifications)
	return notifications
}

// GetNotificationsByType retrieves notifications from the database based on the type
func GetNotificationsByType(notificationType NotificationType) []Notification {
	var notifications []Notification
	store.DB.Where("type = ?", notificationType).Find(&notifications)
	return notifications
}

// GetNotificationsByRecipient retrieves notifications from the database based on the recipient
func GetNotificationsByRecipient(recipient string) []Notification {
	var notifications []Notification
	store.DB.Where("recipient = ?", recipient).Find(&notifications)
	return notifications
}

// GetNotificationsByUser retrieves notifications from the database based on the user
func GetNotificationsByUser(userID uint) []Notification {
	var notifications []Notification
	store.DB.Where("user_id = ?", userID).Find(&notifications)
	return notifications
}
