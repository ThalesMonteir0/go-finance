package routes

import (
	"financial-go/internal/infra/web/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userHandler handlers.UserHandler, fixedAccHandler handlers.FixedAccountHandler) {
	e.POST("/user", userHandler.CreateUser)
	e.DELETE("/user/:cel", userHandler.DeleteUser)

	e.POST("/fixedAccount", fixedAccHandler.CreateFixedAccount)
	e.DELETE("/fixedAccount/:id", fixedAccHandler.DeleteFixedAccount)
	e.PUT("/fixedAccount/paid/:id", fixedAccHandler.PaidFixedAccount)
}
