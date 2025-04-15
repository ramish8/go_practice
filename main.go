package main

import (
	"go_practice/controller"
	"go_practice/infrastructure"
	"go_practice/infrastructure/repository"
	"go_practice/presenter"
	"go_practice/usecase"
	"log"
)

func main() {
	db, err := infrastructure.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	ur := repository.NewUserRepository(db)
	up := presenter.NewUserPresenter()
	ep := presenter.NewErrorPresenter()
	uu := usecase.NewCreateUserUsecase(ur)
	uc := controller.NewUserController(uu, up, ep)
	server := infrastructure.NewServer(uc)
	server.InitRouter()
}
