package storage

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/hamkorbank/go_rest_struct/storage/postgres"
	"gitlab.com/hamkorbank/go_rest_struct/storage/repo"
)

// PostgresStorageI ...
type PostgresStorageI interface {
	Application() repo.ApplicationStorageI
}

type storagePostgres struct {
	db              *sqlx.DB
	applicationRepo repo.ApplicationStorageI
}

// NewStoragePostgres ...
func NewStoragePostgres(db *sqlx.DB) PostgresStorageI {
	return &storagePostgres{
		db:              db,
		applicationRepo: postgres.NewApplicationRepo(db),
	}
}

// Application ...
func (s storagePostgres) Application() repo.ApplicationStorageI {
	return s.applicationRepo
}
