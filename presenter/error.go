package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IErrorPresenter interface {
	BadRequest(ctx echo.Context, err error) error
	InternalServerError(ctx echo.Context, err error) error
}

type ErrorPresenter struct{}

// BadRequest implements IErrorPresenter.
func (p *ErrorPresenter) BadRequest(ctx echo.Context, err error) error {
	panic("unimplemented")
}

// InternalServerError implements IErrorPresenter.
func (p *ErrorPresenter) InternalServerError(ctx echo.Context, err error) error {
	panic("unimplemented")
}

func NewErrorPresenter() *ErrorPresenter {
	return &ErrorPresenter{}
}

func (p *ErrorPresenter) PresentBadRequest(c echo.Context, err error) error {
	response := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	return c.JSON(http.StatusBadRequest, response)
}

func (p *ErrorPresenter) PresentInternalServerError(c echo.Context, err error) error {
	response := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	return c.JSON(http.StatusInternalServerError, response)
}
