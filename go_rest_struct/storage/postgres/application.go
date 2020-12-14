package postgres

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/hamkorbank/go_rest_struct/models"
	"gitlab.com/hamkorbank/go_rest_struct/storage/repo"
)

type applicationRepo struct {
	db *sqlx.DB
}

// NewApplicationRepo ...
func NewApplicationRepo(db *sqlx.DB) repo.ApplicationStorageI {
	return &applicationRepo{db: db}
}

// Get ...
func (r *applicationRepo) Get(id int) (*models.Application, error) {

	var application = models.Application{
		ID: id,
	}

	return &application, nil
}
