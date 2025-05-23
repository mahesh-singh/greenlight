package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies      MoviesModel
	Permissions PermissionModel
	Users       UserModel
	Tokens      TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MoviesModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
	}
}
