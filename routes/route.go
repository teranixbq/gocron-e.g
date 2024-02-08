package routes

import (
	"gocroneg/handler"
	"gocroneg/repository"
	"gocroneg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RouteInit(f *fiber.App, db *gorm.DB,rdb *redis.Client){
	repository := repository.NewRepository(db,rdb)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	f.Post("/register", handler.CreateUser)
	f.Get("/", handler.GetAll)
}