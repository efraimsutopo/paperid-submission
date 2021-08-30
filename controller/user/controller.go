package user

import (
	userService "github.com/efraimsutopo/paperid-submission/service/user"
)

type Controller interface{}

type controller struct {
	service userService.Service
}

func New(service userService.Service) Controller {
	return &controller{
		service,
	}
}
