package messaging

import (
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)

var (
	rabbitMQConn     *amqp.Connection
	rabbitMQConnOnce sync.Once
	rabbitMQConnErr  error
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
	}
}

func connectRabbitMQWithoutErr(){
	connectRabbitMQ()
}
// connectRabbitMQ is now an internal function that is only called once.
func connectRabbitMQ() error {
	connectionString := getRMQConnectionString()
	rabbitMQConn, rabbitMQConnErr = amqp.Dial(connectionString)
	if rabbitMQConnErr != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %v\n", rabbitMQConnErr)
		// Consider logging the error instead of printing.
		return rabbitMQConnErr
	}
	return nil
}

func getRMQConnectionString() string {
	rabbitMQSettings := config.RabbitMQSettings
	return fmt.Sprintf("amqp://%s:%s@%s:%s/",
		rabbitMQSettings.Username,
		rabbitMQSettings.Password,
		rabbitMQSettings.Host,
		rabbitMQSettings.Port)
}

func Setup(){
	if rabbitMQConn == nil || rabbitMQConn.IsClosed() {
		rabbitMQConnOnce.Do(connectRabbitMQWithoutErr)
		if rabbitMQConnErr != nil {
			log.Fatalf("RabbitMQ setup failed': %v", rabbitMQConnErr)
		}
	}
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		failOnError(err, "failed to open connection")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"userCreated", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "failed to declare exchange")

	q, err := ch.QueueDeclare(
		"userCreated:notificationService",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"userCreated", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
}

func startConsuming(channel *amqp.Channel, queue *amqp.Queue){
	msgs, err := channel.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			//handle message created here
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages")
	<-forever
}



// GetRMQConnection lazily initializes and returns the RabbitMQ connection.
// If the connection is closed or not yet established, it tries to reconnect.
func GetRMQConnection() (*amqp.Connection, error) {
    if rabbitMQConn == nil || rabbitMQConn.IsClosed() {
        if err := connectRabbitMQ(); err != nil {
            log.Printf("RabbitMQ connection error: %v", err)
            // Handle reconnection with exponential backoff
            reconnectWithBackoff()
        }
    }
    return rabbitMQConn, nil
}



func reconnectWithBackoff() {
    for i := 0; ; i++ {
        if err := connectRabbitMQ(); err == nil {
            log.Println("RabbitMQ reconnected")
            return
        }
        sleepDuration := time.Second * time.Duration(2^i) // Exponential backoff
        log.Printf("Reconnecting to RabbitMQ after %v", sleepDuration)
        time.Sleep(sleepDuration)
        // Set a max backoff time to avoid very long waits.
        if sleepDuration > time.Minute {
            sleepDuration = time.Minute
        }
    }
}
