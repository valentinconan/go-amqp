package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"go-amqp/src/routes/health"
	"go-amqp/src/amqp"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	healthRouter.Init(router)

	//launch goroutine for http server
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Erreur lors du lancement du serveur Gin: %v", err)
		}
	}()

	log.Print("Initializing consumer")

	//Launch the consumer in a go routine
	go amqp.Consumer()

	var forever chan struct{}
	log.Print("Infinite loop")
	<-forever

}
