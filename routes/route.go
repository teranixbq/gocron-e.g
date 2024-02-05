package routes

import (
	"gocroneg/handler"
	"gocroneg/repository"
	"gocroneg/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteInit(f *fiber.App, db *gorm.DB){
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	f.Post("/register", handler.CreateUser)
	f.Get("/", handler.GetAll)
}