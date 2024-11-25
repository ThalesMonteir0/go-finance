package handlers

import (
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (f *FixedAccountHandler) PaidFixedAccount(c echo.Context) error {
	fixedAccountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid ID")
		return c.JSON(restErr.Code, restErr.Message)
	}

	paidUseCase := usecase.NewPayFixedAccountUseCase(f.repository)

	if err := paidUseCase.Execute(fixedAccountID); err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(http.StatusOK, "fixed account payment successful")
}