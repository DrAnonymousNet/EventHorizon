package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID           uuid.UUID
	Name          string         `gorm:"type:varchar(255)"`             // The name of the user
	Email         string         `gorm:"type:varchar(255);uniqueIndex"` // The user's email address
	PhoneNumber   string         `gorm:"type:varchar(100);uniqueIndex"` // The user's phone number
	Notifications []Notification `gorm:"foreignKey:UserID"`             // User's notifications
	Devices       []Device       `gorm:"foreignKey:UserID"`             // User's devices

}

// Device represents a user's device in the system
type Device struct {
	gorm.Model
	UserID   uint   `gorm:"index"`                         // Foreign key for the User
	DeviceID string `gorm:"type:varchar(255);uniqueIndex"` // The device identifier for push notifications
	User     User   `gorm:"foreignKey:UserID"`             // The associated user
}

func CreateUser(createData map[string]interface{}) {
	// user := &User{
	// 	Name:        createData["username"].(string),
	// 	Email:       createData["email"].(string),
	// 	PhoneNumber: createData["phone_number"].(string),
	// 	UID:         createData["user"].(uuid.UUID),
	// }

}
