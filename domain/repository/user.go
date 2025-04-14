package repository

import (
	"context"
	"go_practice/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entity.User) (int64, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (int64, error)
	Delete(ctx context.Context, id int64) error
}
