package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/efraimsutopo/paperid-submission/config"
	"github.com/efraimsutopo/paperid-submission/config/constant"
	"github.com/efraimsutopo/paperid-submission/model"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
)

func GenerateToken(user model.User) (string, error) {
	claims := structs.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.Get().JWTExpiredDuration).Unix(),
		},
		Token: structs.Token{
			UserID: user.ID,
			Email:  user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.Get().JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ClaimToken(tokenString string) (*structs.TokenClaims, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	jwtToken, err := jwt.ParseWithClaims(
		tokenString,
		&structs.TokenClaims{},
		keyFunc)
	if err != nil {
		return nil, err
	}

	tokenClaim, ok := jwtToken.Claims.(*structs.TokenClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return tokenClaim, nil
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signing method invalid")
	} else if method != jwt.SigningMethodHS256 {
		return nil, fmt.Errorf("signing method invalid")
	}

	return []byte(config.Get().JWTSecret), nil
}

func GetTokenFromContext(ec echo.Context) (structs.Token, error) {
	token, ok := ec.Get(constant.TokenKey).(structs.Token)
	if !ok {
		return structs.Token{}, errors.New("failed to get token")
	}

	return token, nil
}

func GetTokenStringFromContext(ec echo.Context) string {
	tokenString := ec.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	return tokenString
}
