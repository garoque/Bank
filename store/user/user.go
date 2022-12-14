package user

import (
	"Q2Bank/model"
	"Q2Bank/utils/customErr"
	"context"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
)

type Store interface {
	Create(ctx context.Context, user model.User) (string, error)
	ReadByID(ctx context.Context, id string) (*model.User, error)
	ReadByEmail(ctx context.Context, email string) (*model.User, error)
	ReadByCpf(ctx context.Context, cpf string) (*model.User, error)
	ReadByCnpj(ctx context.Context, cnpj string) (*model.User, error)
	UpdateBalance(ctx context.Context, user model.User) error
}

func NewStore(dbConn *sqlx.DB) Store {
	return &storeImpl{dbConn}
}

type storeImpl struct {
	dbConn *sqlx.DB
}

func (s *storeImpl) Create(ctx context.Context, user model.User) (string, error) {
	id := uuid.New().String()

	_, err := s.dbConn.ExecContext(ctx, `
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
	`, id, user.Name, user.Cpf, user.Cnpj, user.Email, user.Balance, user.IsSeller, user.Password)

	if err != nil {
		fmt.Println("error user.store.Create: ", err.Error())
		return "", customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar um usuário")
	}

	return id, nil
}

func (s *storeImpl) ReadByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)

	err := s.dbConn.GetContext(ctx, user, `
		SELECT
			id, name, cpf, cnpj, email, balance, is_seller, password, created_at
		FROM users WHERE id = ?
	`, id)

	if err != nil {
		fmt.Println("error user.store.ReadByID: ", err.Error())
		return nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário")
	}

	return user, nil
}

func (s *storeImpl) ReadByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)

	err := s.dbConn.GetContext(ctx, user, `
		SELECT
			id, cpf, cnpj, email, balance, is_seller, password, created_at, updated_at
		FROM users WHERE email = ?
	`, email)

	if err != nil {
		fmt.Println("error user.store.ReadByEmail: ", err.Error())
		return nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário")
	}

	return user, nil
}

func (s *storeImpl) ReadByCpf(ctx context.Context, cpf string) (*model.User, error) {
	user := new(model.User)

	err := s.dbConn.GetContext(ctx, user, `
		SELECT
			id, cpf, cnpj, email, balance, is_seller, password, created_at, updated_at
		FROM users WHERE cpf = ?
	`, cpf)

	if err != nil {
		fmt.Println("error user.store.ReadByCpf: ", err.Error())
		return nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário")
	}

	return user, nil
}

func (s *storeImpl) ReadByCnpj(ctx context.Context, cnpj string) (*model.User, error) {
	user := new(model.User)

	err := s.dbConn.GetContext(ctx, user, `
		SELECT
			id, cpf, cnpj, email, balance, is_seller, password, created_at, updated_at
		FROM users WHERE cnpj = ?
	`, cnpj)

	if err != nil {
		fmt.Println("error user.store.ReadByCnpj: ", err.Error())
		return nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário")
	}

	return user, nil
}

func (s *storeImpl) UpdateBalance(ctx context.Context, user model.User) error {
	_, err := s.dbConn.ExecContext(ctx, `
		UPDATE users SET
			balance = ?
		WHERE id = ?
	`, user.Balance, user.ID)

	if err != nil {
		fmt.Println("error user.store.UpdateBalance: ", err.Error())
		return customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira")
	}

	return nil
}
