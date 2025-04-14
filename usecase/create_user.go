package usecase

import (
	"context"
	"go_practice/domain/entity"
	"go_practice/domain/repository"
)

type ICreateUserUsecase interface {
	Execute(ctx context.Context, name string) (int64, error)
}

type createUserUsecase struct {
	ur repository.IUserRepository
}

func NewCreateUserUsecase(ur repository.IUserRepository) *createUserUsecase {
	return &createUserUsecase{ur: ur}
}

func (u *createUserUsecase) Execute(ctx context.Context, name string) (int64, error) {
	user := entity.User{
		Name: name,
	}

	id, err := u.ur.Create(ctx, &user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
