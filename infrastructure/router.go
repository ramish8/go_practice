package infrastructure

import (
	"go_practice/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userController *controller.UserController
}

func NewServer(userController *controller.UserController) *Server {
	return &Server{userController: userController}
}

func (s *Server) InitRouter() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", s.userController.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
