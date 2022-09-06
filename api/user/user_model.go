package user

import "Q2Bank/model"

type CommonUser struct {
	Name     string `json:"name" validate:"required"`
	Cpf      string `json:"cpf" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SellerUser struct {
	Name     string `json:"name" validate:"required"`
	Cnpj     string `json:"cnpj" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (user *CommonUser) ToUser() *model.User {
	return &model.User{
		Name:     user.Name,
		Cpf:      user.Cpf,
		Email:    user.Email,
		Password: user.Password,
		IsSeller: false,
		Balance:  0,
	}
}

func (user *SellerUser) ToUser() *model.User {
	return &model.User{
		Name:     user.Name,
		Cnpj:     user.Cnpj,
		Email:    user.Email,
		Password: user.Password,
		IsSeller: true,
		Balance:  0,
	}
}
