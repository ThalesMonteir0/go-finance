package handlers

import (
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (f *FixedAccountHandler) DeleteFixedAccount(c echo.Context) error {
	fixedAccountID, err := strconv.Atoi(c.Param("id"))
	userID := int(c.Get("userID").(uint))

	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid id")
		return c.JSON(restErr.Code, NewResponseDataErr(restErr.Message))
	}

	deleteUseCase := usecase.NewDeleteFixedAccountUseCase(f.repository)
	if err := deleteUseCase.Execute(fixedAccountID, userID); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccess("fixed account deleted with success"))
}
