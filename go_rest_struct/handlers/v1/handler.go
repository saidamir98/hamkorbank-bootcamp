package v1

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
	"gitlab.com/hamkorbank/go_rest_struct/storage"
)

// HandlerV1 ...
type HandlerV1 struct {
	cfg             config.Config
	storagePostgres storage.PostgresStorageI
}

// New ...
func New(config config.Config, db *sqlx.DB) *HandlerV1 {
	return &HandlerV1{
		cfg:             config,
		storagePostgres: storage.NewStoragePostgres(db),
	}
}

// Ping ...
func (h *HandlerV1) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GetConfig ...
func (h *HandlerV1) GetConfig(c *gin.Context) {
	c.JSON(200, h.cfg)
}

// GetApplication ...
func (h *HandlerV1) GetApplication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err)
	}
	resp, err := h.storagePostgres.Application().Get(id)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, resp)
}
