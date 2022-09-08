package transaction

import (
	"Q2Bank/model"
	"Q2Bank/services/authorization"
	"Q2Bank/store"
	"context"
	"errors"
	"fmt"
)

type App interface {
	Create(ctx context.Context, transaction model.Transaction) (*model.Transaction, error)
}

func NewApp(stores *store.ContainerStore) App {
	return &appImpl{stores: stores}
}

type appImpl struct {
	stores *store.ContainerStore
}

func (a *appImpl) Create(ctx context.Context, request model.Transaction) (*model.Transaction, error) {
	payer, err := a.stores.User.ReadByID(ctx, request.PayerID)
	if err != nil {
		fmt.Println("app.Create.user.ReadByID.payer: ", err.Error())
		return nil, err
	}

	if payer.IsSeller {
		return nil, errors.New("Lojistas não podem enviar dinheiro")
	}

	if payer.Balance-request.Value < 0 {
		return nil, errors.New("Saldo insuficiente")
	}

	payee, err := a.stores.User.ReadByID(ctx, request.PayeeID)
	if err != nil {
		fmt.Println("app.Create.user.ReadByID.payee: ", err.Error())
		return nil, err
	}

	if isAuthorized := authorization.GetAuthorization(); !isAuthorized {
		fmt.Println("app.Create.authorization.GetAuthorization()")
		return nil, err
	}

	id, err := a.stores.Transaction.Create(ctx, request)
	if err != nil {
		fmt.Println("app.Create.Transaction.Create: ", err.Error())
		return nil, err
	}

	transaction, err := a.stores.Transaction.ReadByID(ctx, id)
	if err != nil {
		fmt.Println("app.Create.Transaction.ReadByID: ", err.Error())
		return nil, err
	}

	payer.Balance -= request.Value
	payee.Balance += request.Value
	// @TODO: alterar status
	return transaction, nil
}
