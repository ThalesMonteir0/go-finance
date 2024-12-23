package main

import (
	"financial-go/internal/infra/database"
	"financial-go/internal/infra/database/config"
	"financial-go/internal/infra/web/handlers"
	"financial-go/internal/infra/web/routes"
	"financial-go/internal/jobs"
	"financial-go/internal/usecase/fixed_account"
	"financial-go/internal/usecase/movement"
	"financial-go/internal/usecase/users"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron"
	"gorm.io/gorm"
	"log"
)

type handlersRoutes struct {
	userHandler         handlers.UserHandler
	movementHandler     handlers.MovementHandler
	fixedAccountHandler handlers.FixedAccountHandler
}

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

	routesHandlers := injectRoutes(db)

	routes.InitRoutes(
		e,
		routesHandlers.userHandler,
		routesHandlers.fixedAccountHandler,
		routesHandlers.movementHandler,
		db)

	c := cron.New()

	c.AddFunc("0 0 1 * *", func() {
		jobs.UpdateFixedAccountPaidDefault(
			routesHandlers.
				fixedAccountHandler.
				FixedAccountUseCase)
	})

	c.Start()

	if err = e.Start(":5005"); err != nil {
		log.Fatal("err start server")
	}
}

func injectRoutes(db *gorm.DB) handlersRoutes {
	userRepo := database.NewUserRepository(db)
	movementRepo := database.NewMovementRepository(db)
	fixedAccountRepo := database.NewFixedAccountRepository(db)

	userUseCase := users.NewUserUseCase(userRepo)
	movementUseCase := movement.NewMovementUseCase(movementRepo)
	fixedAccountUseCase := fixed_account.NewFixedAccount(fixedAccountRepo)

	userHandler := handlers.NewUserHandler(userUseCase)
	movementHandler := handlers.NewMovementHandler(movementUseCase)
	fixedAccountHandler := handlers.NewFixedAccountHandler(fixedAccountUseCase)

	return handlersRoutes{
		userHandler:         userHandler,
		movementHandler:     movementHandler,
		fixedAccountHandler: fixedAccountHandler,
	}
}
