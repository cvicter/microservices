package queue

import (
	"os"

	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	dsn := "amqp://" +
		os.Getenv("RABBITMQ_DEFAULT_USER") + ":" +
		os.Getenv("RABBITMQ_DEFAULT_PASS") + "@" +
		os.Getenv("RABBITMQ_DEFAULT_HOST") + ":" +
		os.Getenv("RABBITMQ_DEFAULT_PORT") +
		os.Getenv("RABBITMQ_DEFAULT_VHOST")

	conn, _ := amqp.Dial(dsn)

	channel, _ := conn.Channel()

	return channel
}

func Notify(payload []byte, exchange string, routingKey string, ch *amqp.Channel) {
	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}
}

func StartConsuming(ch *amqp.Channel, in chan []byte) {
	q, _ := ch.QueueDeclare(
		os.Getenv("RABBITMQ_CONSUMER_QUEUE"),
		true,
		false,
		false,
		false,
		nil,
	)
	msgs, _ := ch.Consume(
		q.Name,
		"checkout",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for m := range msgs {
			in <- []byte(m.Body)
		}
		close(in)
	}()
}
