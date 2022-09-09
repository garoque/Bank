package transaction

import (
	"Q2Bank/model"
	"Q2Bank/test"
	customErr "Q2Bank/utils/err"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func TestCreate(t *testing.T) {
	exampleRequest := model.Transaction{
		Value:   200,
		PayerID: "payer-id",
		PayeeID: "payee-id",
	}

	query := `
		INSERT INTO transactions (
			id,
			id_payer,
			id_payee,
			value
		) VALUES (?, ?, ?, ?)
	`

	cases := map[string]struct {
		Request     model.Transaction
		ExpectedErr error
		prepare     func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:     exampleRequest,
			ExpectedErr: nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), exampleRequest.PayerID, exampleRequest.PayeeID, exampleRequest.Value).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"deve retornar erro": {
			Request:     exampleRequest,
			ExpectedErr: customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar uma transação"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), exampleRequest.PayerID, exampleRequest.PayeeID, exampleRequest.Value).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar uma transação"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			dbConn, mock := test.GetDB()

			cs.prepare(mock)

			store := NewStore(dbConn)
			id, err := store.Create(context.Background(), cs.Request)

			if err == nil && id == "" {
				t.Error(id)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestReadByID(t *testing.T) {
	id := "id"
	createdAt := time.Now()

	query := `
		SELECT
			id, id_payer, id_payee, value, created_at, updated_at
		FROM transactions WHERE id = ?
	`

	expectedResponse := model.Transaction{
		ID:        id,
		PayerID:   "payer-id",
		PayeeID:   "payee-id",
		Value:     200,
		CreatedAt: createdAt,
		UpdatedAt: nil,
	}

	cases := map[string]struct {
		Request          string
		ExpectedErr      error
		ExpectedResponse *model.Transaction
		prepare          func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:          id,
			ExpectedResponse: &expectedResponse,
			ExpectedErr:      nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(id).
					WillReturnRows(
						test.NewRows("id", "id_payer", "id_payee", "value", "created_at", "updated_at").
							AddRow(id, "payer-id", "payee-id", 200, createdAt, nil),
					)
			},
		},
		"deve retornar erro": {
			Request:          id,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler uma transação"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(id).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler uma transação"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			dbConn, mock := test.GetDB()

			cs.prepare(mock)

			store := NewStore(dbConn)
			user, err := store.ReadByID(context.Background(), cs.Request)

			if diff := cmp.Diff(user, cs.ExpectedResponse); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
