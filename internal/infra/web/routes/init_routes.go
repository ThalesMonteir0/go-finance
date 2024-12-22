package routes

import (
	"financial-go/internal/infra/web/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(
	e *echo.Echo,
	userHandler handlers.UserHandler,
	fixedAccHandler handlers.FixedAccountHandler,
	movementHandler handlers.MovementHandler,
	db *gorm.DB,
) {
	e.POST("/user", userHandler.CreateUser)
	e.DELETE("/user/:cel", userHandler.DeleteUser)

	e.POST("/fixedAccount", handlers.JwtMiddleware(fixedAccHandler.CreateFixedAccount, handlers.SomeErrorHandler, db))
	e.DELETE("/fixedAccount/:id", handlers.JwtMiddleware(fixedAccHandler.DeleteFixedAccount, handlers.SomeErrorHandler, db))
	e.PUT("/fixedAccount/paid/:id", handlers.JwtMiddleware(fixedAccHandler.PaidFixedAccount, handlers.SomeErrorHandler, db))
	e.GET("/fixedAccount", handlers.JwtMiddleware(fixedAccHandler.FindAllFixedAccounts, handlers.SomeErrorHandler, db))

	e.POST("/movements", handlers.JwtMiddleware(movementHandler.CreateMovement, handlers.SomeErrorHandler, db))
	e.GET("/movements", handlers.JwtMiddleware(movementHandler.FindAllMovements, handlers.SomeErrorHandler, db))
	e.DELETE("/movements/:movementID", handlers.JwtMiddleware(movementHandler.DeleteMovement, handlers.SomeErrorHandler, db))
}
