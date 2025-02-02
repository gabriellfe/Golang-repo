package rabbit

import (
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection

func SetupRabbit() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	RabbitConn = conn
	startConsumer()
}

func startConsumer() {
	maxConsumer, _ := strconv.Atoi(os.Getenv("RABBIT_MAX_CONSUMER"))

	for i := 1; i <= maxConsumer; i++ {
		go consume()
	}
}

func consume() {
	ch, _ := RabbitConn.Channel()
	ch.Qos(1, 0, false)
	msgs, err := ch.Consume(
		"go.message.inp",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("message received: %v", string(d.Body))
		}
	}()

	forever <- true
}
