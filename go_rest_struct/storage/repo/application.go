package repo

import (
	"gitlab.com/hamkorbank/go_rest_struct/models"
)

//ApplicationStorageI ...
type ApplicationStorageI interface {
	Get(id int) (*models.Application, error)
}
