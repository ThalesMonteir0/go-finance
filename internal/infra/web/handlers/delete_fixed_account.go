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
	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid id")
		return c.JSON(restErr.Code, restErr.Message)
	}

	deleteUseCase := usecase.NewDeleteFixedAccountUseCase(f.repository)
	if err := deleteUseCase.Execute(fixedAccountID); err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(http.StatusOK, "fixed account deleted with success")
}
