package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
)

// HandlerV1 ...
type HandlerV1 struct {
	cfg config.Config
}

// New ...
func New(config config.Config /*, db *sql.DB*/) *HandlerV1 {
	return &HandlerV1{
		cfg: config,
		// storagePostgres: storage.NewStoragePostgres(db),
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
