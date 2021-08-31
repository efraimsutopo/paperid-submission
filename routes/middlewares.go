package routes

import (
	"net/http"

	"github.com/efraimsutopo/paperid-submission/config/constant"
	"github.com/efraimsutopo/paperid-submission/helper"
	"github.com/labstack/echo/v4"
)

func (r *routes) TokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := helper.GetTokenStringFromContext(c)
			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "empty header Authorization")
			}

			tokenClaim, err := helper.ClaimToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "invalid token")
			}

			c.Set(constant.TokenKey, tokenClaim.Token)

			err = r.Controller.User.CheckValidSession(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "invalid session")
			}

			return next(c)
		}
	}
}
