package main

import (
	"financial-go/internal/infra/database"
	"financial-go/internal/infra/database/config"
	"financial-go/internal/infra/web/handlers"
	"financial-go/internal/infra/web/routes"
	"financial-go/internal/usecase/fixed_account"
	"financial-go/internal/usecase/movement"
	"financial-go/internal/usecase/users"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	db, err := config.OpenConnection()
	if err != nil {
		log.Fatal("database cant running")
	}

	injectAndInitRoutes(e, db)

	if err = e.Start(":5005"); err != nil {
		log.Fatal("err start server")
	}
}

func injectAndInitRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := database.NewUserRepository(db)
	movementRepo := database.NewMovementRepository(db)
	fixedAccountRepo := database.NewFixedAccountRepository(db)

	userUseCase := users.NewUserUseCase(userRepo)
	movementUseCase := movement.NewMovementUseCase(movementRepo)
	fixedAccountUseCase := fixed_account.NewFixedAccount(fixedAccountRepo)

	userHandler := handlers.NewUserHandler(userUseCase)
	movementHandler := handlers.NewMovementHandler(movementUseCase)
	fixedAccountHandler := handlers.NewFixedAccountHandler(fixedAccountUseCase)

	routes.InitRoutes(e, userHandler, fixedAccountHandler, movementHandler, db)
}
