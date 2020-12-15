package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
	v1 "gitlab.com/hamkorbank/go_rest_struct/handler/v1"
	"gitlab.com/hamkorbank/go_rest_struct/pkg/logger"
)

func main() {
	server := gin.Default()

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "project_name")

	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Error("postgres connect error", logger.Error(err))
		return
	}

	handlerV1 := v1.New(cfg, log, db)

	server.GET("/ping", handlerV1.Ping)
	server.GET("/config", handlerV1.GetConfig)
	server.GET("/applications/:id", handlerV1.GetApplication)

	server.Run(cfg.HTTPPort)
}
