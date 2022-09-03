package models

type CommonUser struct {
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
