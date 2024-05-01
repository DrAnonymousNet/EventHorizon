package backends

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)

type FCMNotificationSender struct {

}

func NewFCMNotificationSender() *FCMNotificationSender {
	return &FCMNotificationSender{}
}


func (s *FCMNotificationSender) SendNotification(payload NotificationPayload) error {

	// FCM endpoint for sending messages
	fcmEndpoint := "https://fcm.googleapis.com/fcm/send"
	// Your FCM server key
	serverKey := config.AppSetting.FCMServerKey

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return err
	}

	// Create a new HTTP request to the FCM endpoint
	req, err := http.NewRequest("POST", fcmEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set the appropriate headers
	req.Header.Set("Authorization", "key="+serverKey)
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read and print the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(respBody))
	return nil
}
