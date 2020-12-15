package v1

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"gitlab.com/hamkorbank/go_rest_struct/config"
	"gitlab.com/hamkorbank/go_rest_struct/pkg/logger"
	"gitlab.com/hamkorbank/go_rest_struct/storage"
)

// Handler represents a handler
type Handler struct {
	cfg             config.Config
	log             logger.Logger
	storagePostgres storage.PostgresStorageI
}

// New returns a handler
func New(config config.Config, logger logger.Logger, db *sqlx.DB) *Handler {
	return &Handler{
		cfg:             config,
		log:             logger,
		storagePostgres: storage.NewStoragePostgres(db),
	}
}

// Ping ...
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetConfig ...
func (h *Handler) GetConfig(c *gin.Context) {
	if h.cfg.Environment == "release" {
		h.log.Info("do not return config on release server!", logger.Any("environment", h.cfg.Environment))
		c.JSON(http.StatusOK, h.cfg.Environment)
		return
	}

	c.JSON(http.StatusOK, h.cfg)
}

// GetApplication ...
func (h *Handler) GetApplication(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("error while parsing id param to int", logger.Error(err), logger.Any("id", c.Param("id")))
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	resp, err := h.storagePostgres.Application().Get(id)
	if err != nil {
		h.log.Error("error while getting application data by id from storage", logger.Error(err), logger.Any("id", id))
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, resp)
}
