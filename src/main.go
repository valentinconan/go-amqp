package main

import (
	"log"
	"go-amqp/src/routes"
	"go-amqp/src/amqp"
)

func main() {

	log.Print("Initializing router")
	go router.Init()

	log.Print("Initializing consumer")

	//Launch the consumers
	amqp.Init()

	var forever chan struct{}
	log.Print("Infinite loop")
	<-forever

}
