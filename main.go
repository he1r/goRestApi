package main

import (
	"restApi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	//example get endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// run the backend -- default localhost/8080 port is being used
	r.Run()
}
