package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
)

func main() {
	r := gin.Default()

	cfg := config.Load()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(cfg.HttpPort)
}
