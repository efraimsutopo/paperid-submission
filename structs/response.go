package structs

// General Response
type (
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	MessageResponse struct {
		Message string `json:"message"`
	}
)

// User Response
type (
	UserResponse struct {
		ID    uint64 `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
)

// Session Response
type (
	SessionResponse struct {
		Token string `json:"token"`
	}
)
