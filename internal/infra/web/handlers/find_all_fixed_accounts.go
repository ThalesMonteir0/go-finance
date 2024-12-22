package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (f *FixedAccountHandler) FindAllFixedAccounts(c echo.Context) error {
	userID := int(c.Get("userID").(uint))

	fixedAccounts, err := f.fixedAccountUseCase.FindAllFixedAccounts(userID)
	if err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccessWithData(fixedAccounts, "get all fixed accounts with success"))
}
