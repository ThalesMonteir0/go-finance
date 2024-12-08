package handlers

import (
	"financial-go/internal/entity"
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FixedAccountHandler struct {
	repository entity.FixedAccountInterface
}

func NewFixedAccountHandler(repo entity.FixedAccountInterface) FixedAccountHandler {
	return FixedAccountHandler{
		repository: repo,
	}
}

func (f *FixedAccountHandler) CreateFixedAccount(c echo.Context) error {
	var fixedAccount usecase.FixedAccountDTO
	UserID := c.Get("userID").(uint)
	fixedAccount.UserID = int(UserID)

	if err := c.Bind(&fixedAccount); err != nil {
		errRest := rest_err.NewBadRequestError(err.Error())
		return c.JSON(errRest.Code, NewResponseDataErr(errRest.Message))
	}

	createFixedAccountUseCase := usecase.NewCreateFixedAccount(f.repository)

	if err := createFixedAccountUseCase.Execute(fixedAccount); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusCreated, NewResponseDataSuccess("fixed account created with success"))
}
