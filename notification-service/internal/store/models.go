package store

import (
	"time"
)

type User struct {
	UserID          string
	OrganizationID  string // Optional, for multi-tenancy support
	OrganizationName string
}

type DeliveryMethod struct {
	Email           bool
	SMS             bool
	PushNotification bool
	Webhook         bool
}

type DeliveryDetails struct {
	EmailAddresses  []string
	SenderName      string
	SenderEmail     string
	PhoneNumbers    []string
	DeviceTokens    []string // For push notifications
	WebhookURL      string
	WebhookSecret   string // Optional, for webhook verification
}

type Placeholder struct {
	Key          string
	FallbackValue string // Optional
}

type NotificationContent struct {
	HTMLContent    string
	TextContent    string // Text-only version for email
	SubjectLine    string // For emails and push notifications
	Placeholders   []Placeholder
	CustomBranding bool // Include company logos or branding
}

type TriggerCondition struct {
	EventName      string // e.g., "order-created"
	CustomCriteria map[string]interface{} // Additional custom criteria for triggering
}

type NotificationConfig struct {
	User              User
	NotificationType  string // e.g., "order-created"
	Priority          string // e.g., "high", "medium", "low"
	DeliveryMethods   DeliveryMethod
	DeliveryDetails   DeliveryDetails
	Content           NotificationContent
	TriggerConditions TriggerCondition
	Scheduling        time.Time // Optional, for delayed delivery
	Language          string    // Optional, for localization
	Compliance        struct {
		ConsentObtained bool
		DataHandlingNotes string // Notes on data privacy and handling
	}
	UnsubscribeMechanism bool // Include an unsubscribe mechanism
}
