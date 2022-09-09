package user

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
	exampleRequest := model.User{
		Name:     "User Teste",
		Cpf:      "123.456.789-23",
		Email:    "useremail@gmail.com",
		Password: "password",
	}

	query := `
		INSERT INTO users (
			id,
			name,
			cpf,
			cnpj,
			email,
			balance,
			is_seller,
			password
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	cases := map[string]struct {
		Request     model.User
		ExpectedErr error
		prepare     func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:     exampleRequest,
			ExpectedErr: nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), exampleRequest.Name, exampleRequest.Cpf, exampleRequest.Cnpj, exampleRequest.Email, exampleRequest.Balance, exampleRequest.IsSeller, exampleRequest.Password).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"deve retornar erro": {
			Request:     exampleRequest,
			ExpectedErr: customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar um usuário"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), exampleRequest.Name, exampleRequest.Cpf, exampleRequest.Cnpj, exampleRequest.Email, exampleRequest.Balance, exampleRequest.IsSeller, exampleRequest.Password).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar um usuário"))
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
			id, name, cpf, cnpj, email, balance, is_seller, password, created_at
		FROM users WHERE id = ?
	`

	expectedResponse := model.User{
		ID:        id,
		Name:      "User",
		Cpf:       "123.456.789-10",
		Cnpj:      "",
		Email:     "email@gmail.com",
		Balance:   0,
		IsSeller:  true,
		Password:  "password",
		CreatedAt: createdAt,
	}

	cases := map[string]struct {
		Request          string
		ExpectedErr      error
		ExpectedResponse *model.User
		prepare          func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:          id,
			ExpectedResponse: &expectedResponse,
			ExpectedErr:      nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(id).
					WillReturnRows(
						test.NewRows("id", "name", "cpf", "cnpj", "email", "balance", "is_seller", "password", "created_at").
							AddRow(id, "User", "123.456.789-10", "", "email@gmail.com", 0, true, "password", createdAt),
					)
			},
		},
		"deve retornar erro": {
			Request:          id,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(id).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
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

func TestReadByEmail(t *testing.T) {
	email := "email"
	createdAt := time.Now()

	query := `
		SELECT
			id, cpf, cnpj, email, balance, is_seller, password, created_at, updated_at
		FROM users WHERE email = ?
	`

	expectedResponse := model.User{
		ID:        "id",
		Name:      "User",
		Cpf:       "123.456.789-10",
		Cnpj:      "",
		Email:     email,
		Balance:   0,
		IsSeller:  true,
		Password:  "password",
		CreatedAt: createdAt,
	}

	cases := map[string]struct {
		Request          string
		ExpectedErr      error
		ExpectedResponse *model.User
		prepare          func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:          email,
			ExpectedResponse: &expectedResponse,
			ExpectedErr:      nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(email).
					WillReturnRows(
						test.NewRows("id", "name", "cpf", "cnpj", "email", "balance", "is_seller", "password", "created_at").
							AddRow("id", "User", "123.456.789-10", "", email, 0, true, "password", createdAt),
					)
			},
		},
		"deve retornar erro": {
			Request:          email,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(email).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			dbConn, mock := test.GetDB()

			cs.prepare(mock)

			store := NewStore(dbConn)
			user, err := store.ReadByEmail(context.Background(), cs.Request)

			if diff := cmp.Diff(user, cs.ExpectedResponse); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestReadByCpf(t *testing.T) {
	cpf := "cpf"
	createdAt := time.Now()

	query := `
		SELECT
			id, cpf, cnpj, email, balance, is_seller, password, created_at, updated_at
		FROM users WHERE cpf = ?
	`

	expectedResponse := model.User{
		ID:        "id",
		Name:      "User",
		Cpf:       cpf,
		Cnpj:      "",
		Email:     "email",
		Balance:   0,
		IsSeller:  true,
		Password:  "password",
		CreatedAt: createdAt,
	}

	cases := map[string]struct {
		Request          string
		ExpectedErr      error
		ExpectedResponse *model.User
		prepare          func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:          cpf,
			ExpectedResponse: &expectedResponse,
			ExpectedErr:      nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(cpf).
					WillReturnRows(
						test.NewRows("id", "name", "cpf", "cnpj", "email", "balance", "is_seller", "password", "created_at").
							AddRow("id", "User", cpf, "", "email", 0, true, "password", createdAt),
					)
			},
		},
		"deve retornar erro": {
			Request:          cpf,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(cpf).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			dbConn, mock := test.GetDB()

			cs.prepare(mock)

			store := NewStore(dbConn)
			user, err := store.ReadByCpf(context.Background(), cs.Request)

			if diff := cmp.Diff(user, cs.ExpectedResponse); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestReadByCnpj(t *testing.T) {
	cnpj := "cnpj"
	createdAt := time.Now()

	query := `
		SELECT
			id, cpf, cnpj, email, balance, is_seller, password, created_at, updated_at
		FROM users WHERE cnpj = ?
	`

	expectedResponse := model.User{
		ID:        "id",
		Name:      "User",
		Cpf:       "",
		Cnpj:      cnpj,
		Email:     "email",
		Balance:   0,
		IsSeller:  true,
		Password:  "password",
		CreatedAt: createdAt,
	}

	cases := map[string]struct {
		Request          string
		ExpectedErr      error
		ExpectedResponse *model.User
		prepare          func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:          cnpj,
			ExpectedResponse: &expectedResponse,
			ExpectedErr:      nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(cnpj).
					WillReturnRows(
						test.NewRows("id", "name", "cpf", "cnpj", "email", "balance", "is_seller", "password", "created_at").
							AddRow("id", "User", "", cnpj, "email", 0, true, "password", createdAt),
					)
			},
		},
		"deve retornar erro": {
			Request:          cnpj,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(cnpj).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			dbConn, mock := test.GetDB()

			cs.prepare(mock)

			store := NewStore(dbConn)
			user, err := store.ReadByCnpj(context.Background(), cs.Request)

			if diff := cmp.Diff(user, cs.ExpectedResponse); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestUpdateBalance(t *testing.T) {
	user := model.User{
		Balance: 200,
		ID:      "user-id",
	}

	query := `
		UPDATE users SET
			balance = ?
		WHERE id = ?
	`

	cases := map[string]struct {
		Request     model.User
		ExpectedErr error
		prepare     func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			Request:     user,
			ExpectedErr: nil,
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(user.Balance, user.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"deve retornar erro": {
			Request:     user,
			ExpectedErr: customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"),
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(user.Balance, user.ID).
					WillReturnError(customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			dbConn, mock := test.GetDB()

			cs.prepare(mock)

			store := NewStore(dbConn)
			err := store.UpdateBalance(context.Background(), cs.Request)

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
