package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
	v1 "gitlab.com/hamkorbank/go_rest_struct/handlers/v1"
)

func main() {
	r := gin.Default()

	cfg := config.Load()

	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	db, err := sqlx.Connect("postgres", psqlConnString)

	if err != nil {
		fmt.Println(err)
		// log.Error("postgres connect error",
		// 	logger.Error(err))
		return
	}

	handlerV1 := v1.New(cfg, db)

	r.GET("/ping", handlerV1.Ping)
	r.GET("/config", handlerV1.GetConfig)
	r.GET("/applications/:id", handlerV1.GetApplication)

	r.Run(cfg.HTTPPort)
}
