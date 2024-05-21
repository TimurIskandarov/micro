package user

import (
	"context"

	"user/internal/model"
)
func (r *repository) Profile(ctx context.Context, email string) (*model.User, error) {
	return &model.User{}, nil
}