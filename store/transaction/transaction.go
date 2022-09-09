package transaction

import (
	"Q2Bank/model"
	"Q2Bank/utils/customErr"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(ctx context.Context, transaction model.Transaction) (string, error)
	ReadByID(ctx context.Context, id string) (*model.Transaction, error)
}

func NewStore(dbConn *sqlx.DB) Store {
	return &storeImpl{dbConn}
}

type storeImpl struct {
	dbConn *sqlx.DB
}

func (s *storeImpl) Create(ctx context.Context, transaction model.Transaction) (string, error) {
	id := uuid.New().String()

	_, err := s.dbConn.ExecContext(ctx, `
		INSERT INTO transactions (
			id,
			id_payer,
			id_payee,
			value
		) VALUES (?, ?, ?, ?)
	`, id, transaction.PayerID, transaction.PayeeID, transaction.Value)

	if err != nil {
		fmt.Println("error transaction.store.Create: ", err.Error())
		return "", customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar uma transação")
	}

	return id, nil
}

func (s *storeImpl) ReadByID(ctx context.Context, id string) (*model.Transaction, error) {
	transaction := new(model.Transaction)

	err := s.dbConn.GetContext(ctx, transaction, `
		SELECT
			id, id_payer, id_payee, value, created_at, updated_at
		FROM transactions WHERE id = ?
	`, id)

	if err != nil {
		fmt.Println("error transaction.store.ReadByID: ", err.Error())
		return nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler uma transação")
	}

	return transaction, nil
}
