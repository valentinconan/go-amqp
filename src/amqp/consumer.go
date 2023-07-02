package amqp

import (
	"log"
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


	var forever chan struct{}

	// Launch goroutine in order to proceed AMQP messages
	go func() {
		log.Print("Waiting for message to consume")
		for msg := range msgs {
			// process message
			log.Printf("Message received: %s", msg.Body)

			// Mark as read
			err := msg.Ack(false)
			if err != nil {
				log.Printf("Error while acknoledgement: %v", err)
			}
		}
	}()

	<-forever

	log.Println("End of consumer")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}