package main

import (
	"gocroneg/config"
	"gocroneg/routes"
	"gocroneg/scheduler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	f := fiber.New()

	db := config.InitPostgresDB()
	routes.RouteInit(f, db)
	go scheduler.Scheduler(db)
	
	f.Use(cors.New())

	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	f.Listen(":8080")
}
