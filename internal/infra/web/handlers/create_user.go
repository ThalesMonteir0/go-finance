package handlers

import (
	"financial-go/internal/entity"
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userRepository entity.UserRepositoryInterface
}

func NewUserHandler(userRepo entity.UserRepositoryInterface) UserHandler {
	return UserHandler{
		userRepository: userRepo,
	}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user usecase.UserDTO
	if err := c.Bind(&user); err != nil {
		errRest := rest_err.NewBadRequestError(err.Error())
		return c.JSON(errRest.Code, echo.Map{"error": errRest.Message})
	}

	createUserUseCase := usecase.NewCreateUserUseCase(u.userRepository)

	if err := createUserUseCase.Execute(user); err != nil {
		return c.JSON(err.Code, echo.Map{"error": err.Message})
	}

	return c.JSON(http.StatusCreated, echo.Map{"success": "user created!"})
}
