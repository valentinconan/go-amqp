package amqp

import (
	"log"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer() {

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"sample.message.send",    // name
		true, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		amqp.Table{"x-queue-mode": "lazy"},   // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
	}


	log.Println("Connection to queue")
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")


	for {
		select {
		case delivery := <-msgs:
			// Ack a message every 2 seconds
			log.Printf("Received message: %s\n", delivery.Body)
			if err := delivery.Ack(false); err != nil {
				log.Printf("Error acknowledging message: %s\n", err)
			}
			<-time.After(time.Second * 2)
		}
	}


	log.Println("End of consumer")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}