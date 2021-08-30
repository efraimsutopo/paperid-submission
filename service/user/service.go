package user

import (
	userRepository "github.com/efraimsutopo/paperid-submission/repository/user"
)

type Service interface {
}

type service struct {
	userRepository userRepository.Repository
}

func New(
	userRepository userRepository.Repository,
) Service {
	return &service{
		userRepository,
	}
}
