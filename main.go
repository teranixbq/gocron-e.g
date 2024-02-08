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
	rdb := config.InitRedis()
	routes.RouteInit(f, db, rdb)
	go scheduler.Scheduler(db,rdb)

	f.Use(cors.New())

	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	f.Listen(":8080")
}
