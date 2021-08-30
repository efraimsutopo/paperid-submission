package structs

// Transaction
type (
	CreateTransactionRequest struct {
		Type   string `json:"type" validate:"required,oneof=income expense"`
		Amount uint64 `json:"amount" validate:"required"`
		Note   string `json:"note"`
	}

	UpdateTransactionRequest struct {
		ID     uint64  `json:"-"`
		Type   *string `json:"type" validate:"omitempty,oneof=income expense"`
		Amount *uint64 `json:"amount"`
		Note   *string `json:"note"`
	}
)
