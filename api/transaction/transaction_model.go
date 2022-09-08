package transaction

type RequestTransaction struct {
	Value   float64 `json:"value" validate:"required"`
	PayerID string  `json:"payerId" validate:"required"`
	PayeeID string  `json:"payeeId" validate:"required"`
}
