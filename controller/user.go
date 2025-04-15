package controller

import (
	"go_practice/api"
	"go_practice/presenter"
	"go_practice/usecase"
	"go_practice/usecase/dto/input"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	createUserUsecase usecase.ICreateUserUsecase
	userPresenter     presenter.IUserPresenter
	errorPresenter    presenter.IErrorPresenter
}

func NewUserController(
	createUserUsecase usecase.ICreateUserUsecase,
	userPresenter presenter.IUserPresenter,
	errorPresenter presenter.IErrorPresenter,
) *UserController {
	return &UserController{
		createUserUsecase: createUserUsecase,
		userPresenter:     userPresenter,
		errorPresenter:    errorPresenter,
	}
}

func (c *UserController) CreateUser(e echo.Context) error {
	param := api.CreateUserRequest{}
	if err := e.Bind(&param); err != nil {
		return c.errorPresenter.BadRequest(e, err)
	}

	ctx := e.Request().Context()
	input := &input.CreateUserInput{
		Name: param.Name,
	}

	user, err := c.createUserUsecase.Execute(ctx, input)

	if err != nil {
		return c.errorPresenter.InternalServerError(e, err)
	}

	return c.userPresenter.CreateUser(e, user)
}
