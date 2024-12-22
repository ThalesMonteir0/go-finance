package handlers

import (
	"financial-go/internal/usecase/users"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserUseCase users.UserUseCaseInterface
}

func NewUserHandler(userUseCase users.UserUseCaseInterface) UserHandler {
	return UserHandler{
		UserUseCase: userUseCase,
	}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user users.UserDTO
	if err := c.Bind(&user); err != nil {
		errRest := rest_err.NewBadRequestError(err.Error())
		return c.JSON(errRest.Code, NewResponseDataErr(errRest.Message))
	}

	if err := u.UserUseCase.Create(user); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusCreated, NewResponseDataSuccess("user created!"))
}
