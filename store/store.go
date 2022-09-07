package store

import (
	"Q2Bank/store/transaction"
	"Q2Bank/store/user"

	"github.com/jmoiron/sqlx"
)

// ContainerStore is a struct to export instances of database
type ContainerStore struct {
	User        user.Store
	Transaction transaction.Store
}

// NewContainerStore create a new instance of repositories database
func NewContainerStore(dbConnection *sqlx.DB) *ContainerStore {
	return &ContainerStore{
		User:        user.NewStore(dbConnection),
		Transaction: transaction.NewStore(dbConnection),
	}
}
