package model

import "time"

type User struct {
	Name      string     `json:"name"`
	Cpf       string     `json:"cpf"`
	Cnpj      string     `json:"cnpj"`
	Email     string     `json:"email"`
	Balance   float64    `json:"balance"`
	IsSeller  bool       `json:"isSeller"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
