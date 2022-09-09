package user

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

	exampleRequest := model.User{
		Name:     "User Teste",
		Cpf:      "123.456.789-23",
		Email:    "useremail@gmail.com",
		Password: "password",
	}

	exampleRequestSeller := model.User{
		Name:     "User Teste",
		Cnpj:     "123.456.789-23",
		Email:    "useremail@gmail.com",
		Password: "password",
	}

	exampleResponse := model.User{
		ID:        "id",
		Name:      "User Teste",
		Cpf:       "123.456.789-23",
		Cnpj:      "",
		Email:     "useremail@gmail.com",
		Balance:   0,
		IsSeller:  true,
		Password:  "password",
		CreatedAt: createdAt,
	}

	exampleResponseSeller := model.User{
		ID:        "id",
		Name:      "User Teste",
		Cnpj:      "123.456.789-23",
		Cpf:       "",
		Email:     "useremail@gmail.com",
		Balance:   0,
		IsSeller:  true,
		Password:  "password",
		CreatedAt: createdAt,
	}

	cases := map[string]struct {
		Request          model.User
		ExpectedResponse *model.User
		ExpectedErr      error
		prepare          func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore)
	}{
		"deve retornar sucesso": {
			Request:          exampleRequest,
			ExpectedResponse: &exampleResponse,
			ExpectedErr:      nil,
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().ReadByCpf(gomock.Any(), exampleRequest.Cpf).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).
					Return("id", nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), "id").Times(1).
					Return(&exampleResponse, nil)
			},
		},
		"deve retornar sucesso cadastrando vendedor": {
			Request:          exampleRequestSeller,
			ExpectedResponse: &exampleResponseSeller,
			ExpectedErr:      nil,
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequestSeller.Email).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().ReadByCnpj(gomock.Any(), exampleRequestSeller.Cnpj).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).
					Return("id", nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), "id").Times(1).
					Return(&exampleResponseSeller, nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário email": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
		"deve retornar erro: Já existe um usuário com esse email": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusConflict, "Já existe um usuário com esse email"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(&model.User{Email: "useremail@gmail.com"}, nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário cpf": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(nil, nil)

				mockUser.EXPECT().ReadByCpf(gomock.Any(), exampleRequest.Cpf).Times(1).
					Return(nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
		"deve retornar erro: Já existe um usuário com esse cpf": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusConflict, "Já existe um usuário com esse cpf"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(nil, nil)

				mockUser.EXPECT().ReadByCpf(gomock.Any(), exampleRequest.Cpf).Times(1).
					Return(&model.User{Cpf: "123.456.789-23"}, nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário cnpj": {
			Request:          exampleRequestSeller,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequestSeller.Email).Times(1).
					Return(nil, nil)

				mockUser.EXPECT().ReadByCnpj(gomock.Any(), exampleRequestSeller.Cnpj).Times(1).
					Return(nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
		"deve retornar erro: Já existe um usuário com esse cnpj": {
			Request:          exampleRequestSeller,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusConflict, "Já existe um usuário com esse cnpj"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequestSeller.Email).Times(1).
					Return(nil, nil)

				mockUser.EXPECT().ReadByCnpj(gomock.Any(), exampleRequestSeller.Cnpj).Times(1).
					Return(&model.User{Cnpj: "123.456.789-23"}, nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao cadastrar um usuário": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().ReadByCpf(gomock.Any(), exampleRequest.Cpf).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).
					Return("", customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao cadastrar um usuário"))
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário": {
			Request:          exampleRequest,
			ExpectedResponse: nil,
			ExpectedErr:      customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByEmail(gomock.Any(), exampleRequest.Email).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().ReadByCpf(gomock.Any(), exampleRequest.Cpf).Times(1).
					Return(nil, ERROR_NO_ROWS)

				mockUser.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).
					Return("id", nil)

				mockUser.EXPECT().ReadByID(gomock.Any(), "id").Times(1).
					Return(nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
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

func TestIncreaseBalance(t *testing.T) {
	var value float64 = 200
	userId := "id"

	cases := map[string]struct {
		RequestValue  float64
		RequestUserID string
		ExpectedErr   error
		prepare       func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore)
	}{
		"deve retornar sucesso": {
			RequestValue:  value,
			RequestUserID: userId,
			ExpectedErr:   nil,
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), userId).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 500}).Times(1).Return(nil)
			},
		},
		"deve retornar erro: Ocorreu um erro ao ler um usuário": {
			RequestValue:  value,
			RequestUserID: userId,
			ExpectedErr:   customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), userId).Times(1).Return(nil, customErr.New(http.StatusInternalServerError, "Ocorreu um erro ao ler um usuário"))
			},
		},
		"deve retornar erro: Ocorreu um erro atualizar a carteira": {
			RequestValue:  value,
			RequestUserID: userId,
			ExpectedErr:   customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"),
			prepare: func(mockUser *mocks.MockUserStore, mockTransaction *mocks.MockTRansactionStore) {
				mockUser.EXPECT().ReadByID(gomock.Any(), userId).Times(1).Return(&model.User{
					IsSeller: false,
					Balance:  300,
				}, nil)

				mockUser.EXPECT().UpdateBalance(gomock.Any(), model.User{Balance: 500}).Times(1).Return(customErr.New(http.StatusInternalServerError, "Ocorreu um erro atualizar a carteira"))
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

			err := app.IncreaseBalance(ctx, cs.RequestValue, cs.RequestUserID)

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
