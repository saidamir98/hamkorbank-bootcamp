package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
	v1 "gitlab.com/hamkorbank/go_rest_struct/handlers/v1"
)

func main() {
	r := gin.Default()

	cfg := config.Load()

	handlerV1 := v1.New(cfg)

	r.GET("/ping", handlerV1.Ping)
	r.GET("/config", handlerV1.GetConfig)

	r.Run(cfg.HttpPort)
}
