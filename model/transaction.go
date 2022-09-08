package model

import "time"

type Transaction struct {
	ID        string     `json:"id"`
	Value     float64    `json:"value"`
	PayerID   string     `json:"payerId" db:"id_payer"`
	PayeeID   string     `json:"payeeId" db:"id_payee"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
