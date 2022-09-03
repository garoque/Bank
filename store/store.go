package store

import (
	"github.com/jmoiron/sqlx"
)

// ContainerStore is a struct to export instances of database
type ContainerStore struct {
}

// NewContainerStore create a new instance of repositories database
func NewContainerStore(dbConnection *sqlx.DB) *ContainerStore {
	return &ContainerStore{}
}
