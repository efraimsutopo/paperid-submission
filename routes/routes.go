package routes

import (
	"github.com/efraimsutopo/paperid-submission/controller"
	"github.com/labstack/echo/v4"
)

type Routes interface {
	RegisterRoutes(ec *echo.Echo)
}

type routes struct {
	controller.Controller
}

func New(ctrl controller.Controller) Routes {
	return &routes{
		Controller: ctrl,
	}
}

func (r *routes) RegisterRoutes(ec *echo.Echo) {

	user := ec.Group("/user")
	user.POST("/register", r.Controller.User.Register)
	user.POST("/login", r.Controller.User.Login)
	user.POST("/logout", r.Controller.User.Logout, r.TokenMiddleware())
	user.GET("/profile", r.Controller.User.Profile, r.TokenMiddleware())

	transaction := ec.Group("/transaction", r.TokenMiddleware())
	transaction.GET("", r.Controller.Transaction.GetAllInPagination)
	transaction.GET("/summary", r.Controller.Transaction.GetTransactionSummary)
	transaction.GET("/:transactionID", r.Controller.Transaction.GetTransactionByID)
	transaction.POST("", r.Controller.Transaction.CreateTransaction)
	transaction.PUT("/:transactionID", r.Controller.Transaction.UpdateTransactionByID)
	transaction.DELETE("/:transactionID", r.Transaction.DeleteTransactionByID)

}
