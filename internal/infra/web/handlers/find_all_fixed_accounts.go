package handlers

import (
	"financial-go/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (f *FixedAccountHandler) FindAllFixedAccounts(c echo.Context) error {
	userID := int(c.Get("userID").(uint))

	findFixedAccountsUseCase := usecase.NewFindAllFixedAccountsUsecase(f.repository)
	fixedAccounts, err := findFixedAccountsUseCase.Execute(userID)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(http.StatusOK, fixedAccounts)

}
