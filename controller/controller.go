package controller

import (
	transactionController "github.com/efraimsutopo/paperid-submission/controller/transaction"
	userController "github.com/efraimsutopo/paperid-submission/controller/user"
)

type Controller struct {
	User        userController.Controller
	Transaction transactionController.Controller
}

func New(
	user userController.Controller,
	transaction transactionController.Controller,
) Controller {
	return Controller{
		user,
		transaction,
	}
}
