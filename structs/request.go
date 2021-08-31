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

// User
type (
	RegisterUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=32"`
		Name     string `json:"name" validate:"required"`
	}
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}
)
