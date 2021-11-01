package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var settings Settings

func init() {
	settings.Load()
}

func main() {
	router := gin.Default()
	router.POST("/redirection/:network/:source/:destination", func(c *gin.Context) {
		r := Redirection{
			Source:      c.Param("source"),
			Destination: c.Param("destination"),
			Network:     c.Param("network"),
		}
		settings.AddRedirection(r)

		go r.Listen()

		c.Status(200)
	})

	router.GET("/redirection", func(context *gin.Context) {
		context.JSON(200, settings.Redirections)
	})

	// ToDo: Add a DELETE to remove an active redir

	for _, redirection := range settings.Redirections {
		go func(r Redirection) {
			r.Listen()
		}(redirection)
	}

	err := router.Run(settings.WebAddress)
	if err != nil {
		log.Fatalln(err)
	}
}
