package handlers

import (
	"financial-go/internal/entity"
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserRepository entity.UserRepositoryInterface
}

func NewUserHandler(userRepo entity.UserRepositoryInterface) UserHandler {
	return UserHandler{
		UserRepository: userRepo,
	}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user usecase.UserDTO
	if err := c.Bind(&user); err != nil {
		errRest := rest_err.NewBadRequestError(err.Error())
		return c.JSON(errRest.Code, NewResponseDataErr(errRest.Message))
	}

	createUserUseCase := usecase.NewCreateUserUseCase(u.UserRepository)

	if err := createUserUseCase.Execute(user); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusCreated, NewResponseDataSuccess("user created!"))
}
