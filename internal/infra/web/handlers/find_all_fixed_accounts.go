package handlers

import (
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (f *FixedAccountHandler) FindAllFixedAccounts(c echo.Context) error {
	userID, errConv := strconv.Atoi(c.Param("userID"))
	if errConv != nil {
		restErr := rest_err.NewBadRequestError("invalid id")
		return c.JSON(restErr.Code, restErr.Message)
	}

	findFixedAccountsUseCase := usecase.NewFindAllFixedAccountsUsecase(f.repository)
	fixedAccounts, err := findFixedAccountsUseCase.Execute(userID)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(http.StatusOK, fixedAccounts)

}
