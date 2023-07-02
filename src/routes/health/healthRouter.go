package healthRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(router *gin.Engine) {


	router.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "UP",
			"version":    "master"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.Status(200)
	})

}
