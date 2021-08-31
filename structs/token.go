package structs

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type (
	TokenClaims struct {
		jwt.StandardClaims
		Token
	}

	Token struct {
		UserID uint64 `json:"userId"`
		Email  string `json:"email"`
	}
)
