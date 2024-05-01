package messaging

import (
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/dranonymousnet/eventhorizon/api/v1/notifier"
	email_channel "github.com/dranonymousnet/eventhorizon/internal/channel/email"
	email_backend "github.com/dranonymousnet/eventhorizon/internal/channel/email/backends"
	"github.com/dranonymousnet/eventhorizon/internal/config"
)

var (
	rabbitMQConn     *amqp.Connection
	rabbitMQConnOnce sync.Once
	rabbitMQConnErr  error
)

// GetRMQConnectionString constructs the RabbitMQ connection string from configuration.
func getRMQConnectionString() string {
	rabbitMQSettings := config.RabbitMQSettings
	return fmt.Sprintf("amqp://%s:%s@%s:%s/",
		rabbitMQSettings.Username,
		rabbitMQSettings.Password,
		rabbitMQSettings.Host,
		rabbitMQSettings.Port)
}

// connectRabbitMQ establishes a new connection to RabbitMQ.
func connectRabbitMQ() error {
	connectionString := getRMQConnectionString()
	rabbitMQConn, rabbitMQConnErr = amqp.Dial(connectionString)
	if rabbitMQConnErr != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %v\n", rabbitMQConnErr)
		return rabbitMQConnErr
	}

	return nil
}

// Setup initializes RabbitMQ by ensuring connection and setting up the necessary queue and exchange.
func Setup() {
	if rabbitMQConn == nil || rabbitMQConn.IsClosed() {
		rabbitMQConnOnce.Do(func() {
			if err := connectRabbitMQ(); err != nil {
				log.Fatalf("RabbitMQ setup failed: %v", err)
			}
		})
		log.Println("RabbitMQ connected")

	}

	ch, err := rabbitMQConn.Channel()
	failOnError(err, "failed to open a channel")
	//defer ch.Close()
	log.Println("Channel created")

	setupExchangeAndQueue(ch)

}

// setupExchangeAndQueue declares the necessary exchange and queue and binds them.
func setupExchangeAndQueue(ch *amqp.Channel) {
	failOnError(ch.ExchangeDeclare(
		"userCreated", "fanout",
		true,  //durable
		false, //autoDelete
		false, //internal
		false,
		nil,
	), "failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"userCreated:notificationService", false, false, true, false, nil,
	)
	failOnError(err, "Failed to declare a queue")
	log.Println("Exchange declared")

	failOnError(ch.QueueBind(
		q.Name, "", "userCreated", false, nil,
	), "Failed to bind a queue")
	log.Println("Queue binded")
	go startConsuming(ch, q.Name)

}

// startConsuming starts consuming messages from the queue.
func startConsuming(channel *amqp.Channel, queueName string) {
	msgs, err := channel.Consume(
		queueName, "", true, false, false, false, nil,
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			handleDelivery(d.Body)
			log.Printf(" [x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages")
	<-make(chan struct{}) // Block forever
}

// handleDelivery processes the received delivery.
func handleDelivery(body []byte) {
	// var message map[string]interface{}

	// // Unmarshal the JSON into the map
	// err := json.Unmarshal(body, &message)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// messageType, ok := message["message_type"]
	// if !ok{
	// 	fmt.Println("message type missing")
	// 	return
	// }
	// switch messageType{
	// case "userCreated":
	// 	handleUser
	// }
	// // Process the message body
}

func DispatchNotification(requestData *notifier.NotifyRequest) {
	notificationChannels := requestData.NotificationChannels
	for _, ch := range notificationChannels {
		switch ch {
		case "email":
			email := email_backend.NewEmail(
				"me",
				requestData.GetRecipients(),
				requestData.Subject,
				requestData.Message,
				[]string{},
				[]string{},
			)
			email_channel.SendEmail(email)

		case "sms":
			fmt.Println("sms")
		case "push_notification":
			fmt.Println("push")
		}
	}
}

// GetRMQConnection provides the RabbitMQ connection, reconnecting if necessary.
func GetRMQConnection() (*amqp.Connection, error) {
	if rabbitMQConn == nil || rabbitMQConn.IsClosed() {
		reconnectWithBackoff()
	}
	return rabbitMQConn, nil
}

// reconnectWithBackoff attempts to reconnect to RabbitMQ with exponential backoff.
func reconnectWithBackoff() {
	for i := 0; ; i++ {
		if err := connectRabbitMQ(); err == nil {
			log.Println("RabbitMQ reconnected")
			return
		}
		sleepDuration := time.Second * time.Duration(2^i)
		log.Printf("Reconnecting to RabbitMQ after %v", sleepDuration)
		time.Sleep(sleepDuration)

		if sleepDuration > time.Minute {
			sleepDuration = time.Minute
		}
	}
}

// failOnError logs and exits if an error is encountered.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func CloseRBMQConn() {

}
