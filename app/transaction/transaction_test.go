package transaction

import (
	"Q2Bank/mocks"
	"Q2Bank/model"
	"Q2Bank/store"
	"Q2Bank/utils/customErr"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestCreate(t *testing.T) {
	createdAt := time.Now()

	exampleRequest := model.Transaction{
		Value:   200,
		PayerID: "payer-id",
		PayeeID: "payee-id",
	}

	exampleResponse := model.Transaction{
		ID:        "id",
		Value:     200,
		PayerID:   "payer-id",
		PayeeID:   "payee-id",
		CreatedAt: createdAt,
	}

	cases := map[string]struct {
		Request          model.Transaction
		ExpectedResponse *model.Transaction
		ExpectedErr      error
		prepare          func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore)
	}{
		"deve retornar sucesso": {
			Request:          exampleRequest,
			ExpectedResponse: &exampleResponse,
			ExpectedErr:      nil,
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayeeID).Times(1).Return(&model.User{
					Balance: 100,
				}, nil)

				mockTransaction.EXPECT().Create(gomock.Any(), exampleRequest).Times(1).Return("id", nil)

				mockTransaction.EXPECT().ReadByID(gomock.Any(), "id").Times(1).Return(&exampleResponse, nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 100}).Times(1).Return(nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 300}).Times(1).Return(nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
		"deve retornar erro: Lojistas não podem enviar dinheiro": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusBadRequest, "Lojistas não podem enviar dinheiro"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: true,
					Balance:  300,
				}, nil)
			},
		},
		"deve retornar erro: Saldo insuficiente": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusBadRequest, "Saldo insuficiente"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  0,
				}, nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário payee": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayeeID).Times(1).Return(nil,
					customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
		"deve retornar erro: Ocorreu um erro ao cadastrar uma transação": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar uma transação"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayeeID).Times(1).Return(&model.User{
					Balance: 100,
				}, nil)

				mockTransaction.EXPECT().Create(gomock.Any(), exampleRequest).Times(1).Return("",
					customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar uma transação"))
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler uma transação": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler uma transação"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayeeID).Times(1).Return(&model.User{
					Balance: 100,
				}, nil)

				mockTransaction.EXPECT().Create(gomock.Any(), exampleRequest).Times(1).Return("id", nil)

				mockTransaction.EXPECT().ReadByID(gomock.Any(), "id").Times(1).Return(nil,
					customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler uma transação"))
			},
		},
		"deve retornar erro: Ocorreu um erro atualizar a carteira": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayeeID).Times(1).Return(&model.User{
					Balance: 100,
				}, nil)

				mockTransaction.EXPECT().Create(gomock.Any(), exampleRequest).Times(1).Return("id", nil)

				mockTransaction.EXPECT().ReadByID(gomock.Any(), "id").Times(1).Return(&exampleResponse, nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 100}).Times(1).
					Return(customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"))
			},
		},
		"deve retornar erro: Ocorreu um erro atualizar a carteira payee": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayerID).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), exampleRequest.PayeeID).Times(1).Return(&model.User{
					Balance: 100,
				}, nil)

				mockTransaction.EXPECT().Create(gomock.Any(), exampleRequest).Times(1).Return("id", nil)

				mockTransaction.EXPECT().ReadByID(gomock.Any(), "id").Times(1).Return(&exampleResponse, nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 100}).Times(1).Return(nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 300}).Times(1).
					Return(customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			controller, ctx := gomock.WithContext(context.Background(), t)

			mockUser := mocks.NewMockUserStore(controller)
			mockTransaction := mocks.NewMockTRansactionStore(controller)

			cs.prepare(mockUser, mockTransaction)

			app := NewApp(&store.ContainerStore{
				User: mockUser, Transaction: mockTransaction,
			})

			transaction, err := app.Create(ctx, cs.Request)

			if diff := cmp.Diff(transaction, cs.ExpectedResponse); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
