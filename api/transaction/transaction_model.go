package transaction

import "Q2Bank/model"

type RequestTransaction struct {
	Value   float64 `json:"value" validate:"required"`
	PayerID string  `json:"payerId" validate:"required"`
	PayeeID string  `json:"payeeId" validate:"required"`
}

func (t *RequestTransaction) ToTransaction() *model.Transaction {
	return &model.Transaction{
		Value:   t.Value,
		PayerID: t.PayerID,
		PayeeID: t.PayeeID,
	}
}
