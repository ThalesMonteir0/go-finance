package routes

import (
	"financial-go/internal/infra/web/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(
	e *echo.Echo,
	userHandler handlers.UserHandler,
	fixedAccHandler handlers.FixedAccountHandler,
	movementHandler handlers.MovementHandler,
) {
	e.POST("/user", userHandler.CreateUser)
	e.DELETE("/user/:cel", userHandler.DeleteUser)

	e.POST("/fixedAccount", handlers.JwtMiddleware(fixedAccHandler.CreateFixedAccount, handlers.SomeErrorHandler, userHandler.UserRepository))
	e.DELETE("/fixedAccount/:id", handlers.JwtMiddleware(fixedAccHandler.DeleteFixedAccount, handlers.SomeErrorHandler, userHandler.UserRepository))
	e.PUT("/fixedAccount/paid/:id", handlers.JwtMiddleware(fixedAccHandler.PaidFixedAccount, handlers.SomeErrorHandler, userHandler.UserRepository))
	e.GET("/fixedAccount/:userID", fixedAccHandler.FindAllFixedAccounts)

	e.POST("/movements", movementHandler.CreateMovement)
	e.GET("/movements", movementHandler.FindAllMovements)
	e.DELETE("/movements/:movementID", movementHandler.DeleteMovement)
}
