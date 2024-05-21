package user

import (
	"context"

	"user/internal/model"
)

func (r *repository) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	return users, nil
}