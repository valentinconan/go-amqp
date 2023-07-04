package router

import (
	"log"
	"github.com/gin-gonic/gin"
	"go-amqp/src/routes/health"
	"go-amqp/src/routes/amqp"
)

func Init() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	healthRouter.Init(router)
	amqpRouter.Init(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur Gin: %v", err)
	}
}
