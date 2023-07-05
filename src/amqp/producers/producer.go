package producers

import (
	"context"
	"log"
	"time"
	utils "go-amqp/src/amqp/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Produce(body string, queue string){
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()


	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    log.Print("---------", queue)
	if queue == "" {
	    queue = "sample.message.send"
	}

    log.Print("!!!!!!!!!", queue)

	err = ch.PublishWithContext(ctx,
		"", // exchange
		queue,     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}