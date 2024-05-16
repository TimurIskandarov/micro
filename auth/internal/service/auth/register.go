package auth

import (
	"context"
	"net/http"

	"auth/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, in model.RegisterIn) model.RegisterOut {
	// проверка существования пользователя
	if _, ok := users[in.Name]; ok {
		s.logger.Error("this user already exists")
		return model.RegisterOut{
			Message: "this user already exists",
			Status:  http.StatusConflict,
		}
	}

	// хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(in.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		s.logger.Error("this user already exists")
		return model.RegisterOut{
			Message: "this user already exists",
			Status:  http.StatusConflict,
		}
	}

	// сохранение пароля
	users[in.Name] = string(hashedPassword)

	return model.RegisterOut{
		Message: "successful user registration",
		Status:  http.StatusOK,
	}
}
