package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yodramuhammad/ticket-booking/config"
	"github.com/yodramuhammad/ticket-booking/db"
	"github.com/yodramuhammad/ticket-booking/handlers"
	"github.com/yodramuhammad/ticket-booking/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
