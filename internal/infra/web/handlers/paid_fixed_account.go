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
	userID := int(c.Get("userID").(uint))

	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid ID")
		return c.JSON(restErr.Code, NewResponseDataErr(restErr.Message))

	}

	paidUseCase := usecase.NewPayFixedAccountUseCase(f.repository)

	if err := paidUseCase.Execute(fixedAccountID, userID); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccess("fixed account payment successful"))
}
