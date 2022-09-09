package user

import (
	"context"
	"errors"
	"fmt"

	"Q2Bank/model"
	"Q2Bank/store"
	"Q2Bank/utils/hash"
)

var ERROR_NO_ROWS error = errors.New("sql: no rows in result set")

type App interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	alreadyExists(ctx context.Context, email, cpf, cpnj string) (bool, error)
	IncreaseBalance(ctx context.Context, value float64, userID string) error
}

func NewApp(stores *store.ContainerStore) App {
	return &appImpl{stores: stores}
}

type appImpl struct {
	stores *store.ContainerStore
}

func (a *appImpl) Create(ctx context.Context, user model.User) (*model.User, error) {
	alreadyExists, err := a.alreadyExists(ctx, user.Email, user.Cpf, user.Cnpj)
	if err != nil {
		fmt.Println("app.Create.alreadyExists: ", err.Error())
		return nil, err
	}

	if alreadyExists {
		fmt.Println("app.Create.alreadyExists.alreadyExists: ", err.Error())
		return nil, err
	}

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		fmt.Println("app.Create.hash.HashPassword: ", err.Error())
		return nil, err
	}

	id, err := a.stores.User.Create(ctx, user)
	if err != nil {
		fmt.Println("app.Create.stores.User.Create: ", err.Error())
		return nil, err
	}

	userCreated, err := a.stores.User.ReadByID(ctx, id)
	if err != nil {
		fmt.Println("app.Create.stores.User.ReadByID: ", err.Error())
		return nil, err
	}

	return userCreated, nil
}

func (a *appImpl) alreadyExists(ctx context.Context, email, cpf, cnpj string) (bool, error) {
	user, err := a.stores.User.ReadByEmail(ctx, email)
	if err != nil && err.Error() != ERROR_NO_ROWS.Error() {
		return true, err
	}

	if user != nil {
		return true, errors.New("Já existe um usuário com esse email")
	}

	if cpf != "" {
		user, err = a.stores.User.ReadByCpf(ctx, cpf)
		if err != nil && err.Error() != ERROR_NO_ROWS.Error() {
			return true, err
		}

		if user != nil {
			return true, errors.New("Já existe um usuário com esse cpf")
		}
	}

	if cnpj != "" {
		user, err = a.stores.User.ReadByCnpj(ctx, cnpj)
		if err != nil && err.Error() != ERROR_NO_ROWS.Error() {
			return true, err
		}

		if user != nil {
			return true, errors.New("Já existe um usuário com esse cnpj")
		}
	}

	return false, nil
}

func (a *appImpl) IncreaseBalance(ctx context.Context, value float64, userID string) error {
	user, err := a.stores.User.ReadByID(ctx, userID)
	if err != nil {
		fmt.Println("app.IncreaseBalance.User.ReadByID: ", err.Error())
		return err
	}

	user.Balance += value
	err = a.stores.User.UpdateBalance(ctx, *user)
	if err != nil {
		fmt.Println("app.IncreaseBalance.User.UpdateBalance: ", err.Error())
		return err
	}

	return nil
}
