package postgres

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/hamkorbank/go_rest_struct/model"
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
func (r *applicationRepo) Get(id int) (*model.Application, error) {
	var application = model.Application{
		ID: id,
	}

	//
	// Here should be code to query data from applications table by id
	//

	return &application, nil
}
