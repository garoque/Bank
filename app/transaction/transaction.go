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
	// verificar se NÃO é um lojista o payer
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

	// remover dinheiro da carteira do payer
	payer.Balance -= request.Value

	// adicionar dinheiro na carteira do payee
	payee, err := a.stores.User.ReadByID(ctx, request.PayeeID)
	if err != nil {
		fmt.Println("app.Create.user.ReadByID.payee: ", err.Error())
		return nil, err
	}

	payee.Balance += request.Value

	// usar o mock da api externa para validar
	if isAuthorized := authorization.GetAuthorization(); !isAuthorized {
		// rollback
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
		fmt.Println("app.Create.Transaction.Create: ", err.Error())
		return nil, err
	}

	// @TODO: alterar status
	return transaction, nil
}
