package controller

import (
	"context"
	"fmt"
	"go_practice/usecase"
	"net/http"
)

type UserController struct {
	createUserUsecase usecase.ICreateUserUsecase
}

func NewUserController(createUserUsecase usecase.ICreateUserUsecase) *UserController {
	return &UserController{createUserUsecase: createUserUsecase}
}

func (c *UserController) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id, err := c.createUserUsecase.Execute(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created: %d", id)))
}
