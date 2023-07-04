package amqpRouter


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func Init(router *gin.Engine) {


	router.POST("/produce", func(c *gin.Context) {

		// get raw data
		jsonData, err := c.GetRawData()
		if err != nil {
			// Handle error
			log.Print("Unable to get body")

		}
		//return data for display
		c.JSON(http.StatusOK, gin.H{
			"data": string(jsonData)})
	})
}
