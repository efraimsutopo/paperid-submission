package structs

// Transaction
type (
	GetAllInPaginationRequest struct {
		Limit  int     `json:"limit" query:"limit"`
		Page   int     `json:"page" query:"page"`
		Sort   string  `json:"sort" query:"sort"`
		Type   *string `json:"type" query:"type" validate:"omitempty,oneof=income expense"`
		Amount *uint64 `json:"amount" query:"amount"`
		Note   *string `json:"note" query:"note"`
	}

	GetTransactionSummary struct {
		StartDate string `json:"startDate" query:"startDate"`
		EndDate   string `json:"endDate" query:"endDate"`
	}

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
