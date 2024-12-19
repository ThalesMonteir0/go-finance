package handlers

import (
	"financial-go/internal/usecase/fixed_account"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FixedAccountHandler struct {
	fixedAccountUseCase fixed_account.FixedAccountUseCaseInterface
}

func NewFixedAccountHandler(useCase fixed_account.FixedAccountUseCaseInterface) FixedAccountHandler {
	return FixedAccountHandler{
		fixedAccountUseCase: useCase,
	}
}

func (f *FixedAccountHandler) CreateFixedAccount(c echo.Context) error {
	var fixedAccount fixed_account.FixedAccountDTO
	UserID := c.Get("userID").(uint)
	fixedAccount.UserID = int(UserID)

	if err := c.Bind(&fixedAccount); err != nil {
		errRest := rest_err.NewBadRequestError(err.Error())
		return c.JSON(errRest.Code, NewResponseDataErr(errRest.Message))
	}

	if err := f.fixedAccountUseCase.CreateFixedAccount(fixedAccount); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusCreated, NewResponseDataSuccess("fixed account created with success"))
}
