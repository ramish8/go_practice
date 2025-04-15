package usecase

import (
	"context"
	"go_practice/domain/entity"
	"go_practice/domain/repository"
	"go_practice/usecase/dto/input"
	"go_practice/usecase/dto/output"
)

type ICreateUserUsecase interface {
	Execute(ctx context.Context, input *input.CreateUserInput) (*output.CreateUserOutput, error)
}

type createUserUsecase struct {
	ur repository.IUserRepository
}

func NewCreateUserUsecase(ur repository.IUserRepository) *createUserUsecase {
	return &createUserUsecase{ur: ur}
}

func (u *createUserUsecase) Execute(ctx context.Context, input *input.CreateUserInput) (*output.CreateUserOutput, error) {
	e, err := entity.NewUser(input.Name)
	if err != nil {
		return nil, err
	}

	id, err := u.ur.Create(ctx, e)
	if err != nil {
		return nil, err
	}

	user, err := u.ur.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &output.CreateUserOutput{
		User: user,
	}, nil
}
