package routes

import (
	"errors"
	"net/http"

	"github.com/efraimsutopo/paperid-submission/config/constant"
	"github.com/efraimsutopo/paperid-submission/helper"
	"github.com/labstack/echo/v4"
)

func TokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := helper.GetTokenStringFromContext(c)
			if tokenString == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "empty header Authorization")
			}

			tokenClaim, err := helper.ClaimToken(tokenString)
			if err != nil {
				err = errors.New("invalid token")
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			c.Set(constant.TokenKey, tokenClaim.Token)

			return next(c)
		}
	}
}
