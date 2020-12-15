package repo

import (
	"gitlab.com/hamkorbank/go_rest_struct/model"
)

//ApplicationStorageI ...
type ApplicationStorageI interface {
	Get(id int) (*model.Application, error)
}
