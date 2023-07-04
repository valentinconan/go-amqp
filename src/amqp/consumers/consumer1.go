package consumers

import (
	"log"
	"time"
	utils "go-amqp/src/amqp/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer1() {

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
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
	utils.FailOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"sample1.message.send",    // name
		true, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		amqp.Table{"x-queue-mode": "lazy"},   // arguments
	)
	if err != nil {
		utils.FailOnError(err, "Failed to declare a queue")
	}


	log.Println("Connection to queue")
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to bind a queue")
	log.Println("Connected to queue "+q.Name)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

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
