package models

type SellerUser struct {
	Name     string `json:"name"`
	Cnpj     string `json:"cnpj"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
