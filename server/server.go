package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func wallStreet(kindReport string) string {
	if kindReport == "stock" {
		return "Stock delivered to the client."
	} else if kindReport == "transactions" {
		return "Transactions delivered to the client."
	} else if kindReport == "trade" {
		return "trade delivered to the client."
	} else {
		return "reports delivered to the client."
	}
}

func main() {
	//conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	conn, err := amqp.Dial("amqp://admin:Password123@159.65.220.217:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_server",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	//funcion que se ejecutan simultaneamente con otras funciones
	go func() {
		for message := range msgs {
			value := string(message.Body)
			failOnError(err, "Failed to convert body to string")

			log.Printf(" [.] wallStreet(%s)", value)
			response := wallStreet(value)

			err = ch.Publish(
				"",
				message.ReplyTo, // routing key
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: message.CorrelationId,
					Body:          []byte(string(response)),
				})
			failOnError(err, "Failed to publish a message")

			message.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting reports requests")
	<-forever
}
