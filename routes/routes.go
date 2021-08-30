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
	user.POST("/register", nil) // TODO: Replace with real controller func
	user.POST("/login", nil)    // TODO: Replace with real controller func
	user.POST("/logout", nil)   // TODO: Replace with real controller func

	transaction := ec.Group("/transaction")
	transaction.GET("", nil)         // TODO: Replace with real controller func
	transaction.GET("/summary", nil) // TODO: Replace with real controller func
	transaction.GET("/:transactionID", r.Controller.Transaction.GetTransactionByID)
	transaction.POST("", r.Controller.Transaction.CreateTransaction)
	transaction.PUT("/:transactionID", r.Controller.Transaction.UpdateTransactionByID)
	transaction.DELETE("/:transactionID", r.Transaction.DeleteTransactionByID)

}