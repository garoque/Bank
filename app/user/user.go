package user

import (
	"context"
	"errors"

	"Q2Bank/model"
	"Q2Bank/store"
	"Q2Bank/utils/hash"
)

type App interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	alreadyExists(ctx context.Context, email, cpf, cpnj string) (bool, error)
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
		return nil, nil
	}

	if alreadyExists {
		return nil, err
	}

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	id, err := a.stores.User.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	userCreated, err := a.stores.User.ReadByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}

func (a *appImpl) alreadyExists(ctx context.Context, email, cpf, cnpj string) (bool, error) {
	user, err := a.stores.User.ReadByEmail(ctx, email)
	if err != nil {
		return true, err
	}

	if user != nil {
		return true, errors.New("Já existe um usuário com esse email")
	}

	user, err = a.stores.User.ReadByCpf(ctx, cpf)
	if err != nil {
		return true, err
	}

	if user != nil {
		return true, errors.New("Já existe um usuário com esse cpf")
	}

	user, err = a.stores.User.ReadByCnpj(ctx, cnpj)
	if err != nil {
		return true, err
	}

	if user != nil {
		return true, errors.New("Já existe um usuário com esse cnpj")
	}

	return false, nil
}
