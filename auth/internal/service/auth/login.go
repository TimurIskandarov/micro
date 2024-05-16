package auth

import (
	"context"
	"net/http"

	"auth/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, in model.AuthorizedIn) model.AuthorizedOut {
	// функция проверки пароля
	comparator := func(hashedPassword string) bool {
		return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(in.Password)) == nil
	}

	// если пользователь не найден или пароль не совпадает
	if hashedPassword, ok := users[in.Name]; !(ok && comparator(hashedPassword)) {
		s.logger.Error("wrong input data")
		return model.AuthorizedOut{
			Message: "wrong input data",
			Status:  http.StatusForbidden,
		}
	}

	// создание токена
	_, token, _ := s.token.Encode(map[string]interface{}{in.Name: in.Password})

	return model.AuthorizedOut{
		AccessToken: token,
		Message:     "successful user authorization",
		Status:      http.StatusOK,
	}
}
