package presenter

import (
	"go_practice/api"
	"go_practice/usecase/dto/output"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserPresenter interface {
	CreateUser(ctx echo.Context, output *output.CreateUserOutput) error
}

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) CreateUser(ctx echo.Context, output *output.CreateUserOutput) error {
	response := api.CreateUserResponse{
		ID:        output.User.ID,
		Name:      output.User.Name,
		CreatedAt: output.User.CreatedAt,
		UpdatedAt: output.User.UpdatedAt,
	}
	return ctx.JSON(http.StatusCreated, response)
}
