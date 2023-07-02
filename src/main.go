package main

import (
	"log"
	"go-amqp/src/routes"
	"go-amqp/src/amqp"
)

func main() {

	log.Print("Initializing router")
	router.Init()

	log.Print("Initializing consumer")

	//Launch the consumer in a go routine
	go amqp.Consumer()
	go amqp.Consumer1()

	var forever chan struct{}
	log.Print("Infinite loop")
	<-forever

}
