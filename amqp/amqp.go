package amqp

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel

const (
	user     = "amqp_user"
	password = "amqp2018"
	host     = "localhost"
	port     = 5672
)

func init() {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	log.Printf("[AMQP] Connecting to RabbitMQ at %s", addr)
	conn, err := amqp.Dial(addr)
	failOnError(err)

	ch, err = conn.Channel()
	failOnError(err)

	ch.Qos(1, 0, false)
	failOnError(err)
}

func SendMessage(queueName string, message []byte) {
	ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
