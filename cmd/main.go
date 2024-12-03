package main

import (
	"financial-go/internal/infra/database"
	"financial-go/internal/infra/database/config"
	"financial-go/internal/infra/web/handlers"
	"financial-go/internal/infra/web/routes"
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

	if err := e.Start(":5000"); err != nil {
		log.Fatal("err start server")
	}
}

func injectAndInitRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := database.NewUserRepository(db)
	movementRepo := database.NewMovementRepository(db)
	fixedAccountRepo := database.NewFixedAccountRepository(db)

	userHandler := handlers.NewUserHandler(userRepo)
	movementHandler := handlers.NewMovementHandler(movementRepo)
	fixedAccountHandler := handlers.NewFixedAccountHandler(fixedAccountRepo)

	routes.InitRoutes(e, userHandler, fixedAccountHandler, movementHandler)
}
