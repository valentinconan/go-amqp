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
    // catch all panic exception and return 500
    router.Use(gin.Recovery())

    //base endpoint not secured
    root := router.Group("/")

	healthRouter.Init(root)

    //secured group
    amqp := router.Group("/amqp")
    amqp.Use(checkAuthorization())
	amqpRouter.Init(amqp)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du lancement du serveur Gin: %v", err)
	}
}

func checkAuthorization() gin.HandlerFunc {
   return func(c *gin.Context) {
      log.Println("All security check must be proceed here")
   }
}