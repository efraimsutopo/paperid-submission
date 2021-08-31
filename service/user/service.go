package user

import (
	"errors"
	"net/http"

	"github.com/efraimsutopo/paperid-submission/helper"
	"github.com/efraimsutopo/paperid-submission/model"
	sessionRepository "github.com/efraimsutopo/paperid-submission/repository/session"
	userRepository "github.com/efraimsutopo/paperid-submission/repository/user"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
	Register(ec echo.Context, req structs.RegisterUserRequest) (*structs.UserResponse, *structs.ErrorResponse)
	Login(ec echo.Context, req structs.LoginRequest) (*structs.SessionResponse, *structs.ErrorResponse)
	Logout(ec echo.Context) *structs.ErrorResponse
	GetUser(ec echo.Context) (*structs.UserResponse, *structs.ErrorResponse)
	CheckValidSession(ec echo.Context) error
}

type service struct {
	userRepository    userRepository.Repository
	sessionRepository sessionRepository.Repository
}

func New(
	userRepository userRepository.Repository,
	sessionRepository sessionRepository.Repository,
) Service {
	return &service{
		userRepository,
		sessionRepository,
	}
}

func (s *service) Register(ec echo.Context, req structs.RegisterUserRequest) (*structs.UserResponse, *structs.ErrorResponse) {
	hashedPassword, err := helper.GeneratePassword(req.Password)
	if err != nil {
		return nil, &structs.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	res, err := s.userRepository.CreateUser(model.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	})
	if err != nil {
		code := http.StatusInternalServerError
		if helper.IsMySQLErrorDuplicate(err) {
			code = http.StatusBadRequest
			err = errors.New("email already registered")
		}

		return nil, &structs.ErrorResponse{
			Code:    code,
			Message: err.Error(),
		}
	}

	return &structs.UserResponse{
		ID:    res.ID,
		Email: res.Email,
		Name:  res.Name,
	}, nil
}

func (s *service) Login(ec echo.Context, req structs.LoginRequest) (*structs.SessionResponse, *structs.ErrorResponse) {
	user, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			err = errors.New("email not found")
		}

		return nil, &structs.ErrorResponse{
			Code:    code,
			Message: err.Error(),
		}
	}

	if !helper.CheckPassword(req.Password, user.Password) {
		return nil, &structs.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid password",
		}
	}

	tokenStr, err := helper.GenerateToken(*user)
	if err != nil {
		return nil, &structs.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
		}
	}

	_, err = s.sessionRepository.CreateSession(model.Session{
		UserID: user.ID,
		Token:  tokenStr,
	})
	if err != nil {
		return nil, &structs.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to create session",
		}
	}

	return &structs.SessionResponse{
		Token: tokenStr,
	}, nil
}

func (s *service) Logout(ec echo.Context) *structs.ErrorResponse {
	tokenString := helper.GetTokenStringFromContext(ec)

	if err := s.sessionRepository.DeleteSessionByToken(tokenString); err != nil {
		return &structs.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *service) GetUser(ec echo.Context) (*structs.UserResponse, *structs.ErrorResponse) {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return nil, &structs.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	user, err := s.userRepository.GetUserByEmail(token.Email)
	if err != nil {
		return nil, &structs.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &structs.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s *service) CheckValidSession(ec echo.Context) error {
	tokenString := helper.GetTokenStringFromContext(ec)
	_, err := s.sessionRepository.GetSessionByToken(tokenString)
	if err != nil {
		return errors.New("invalid session")
	}
	return nil
}
