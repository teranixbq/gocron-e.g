package handler

import (
	"gocroneg/dto"
	"gocroneg/service"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *handler{
	return &handler{
		service: service,
	}
}

func (eg *handler) CreateUser(f *fiber.Ctx) error{

	input := dto.Request{}

	err  := f.BodyParser(&input)
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message":err.Error(),
		})
	} 

	request := dto.RequesToModel(input)
	err = eg.service.Insert(request)
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return f.Status(201).JSON(fiber.Map{
		"message":"success insert",
	})
}

func (eg *handler) GetAll(f *fiber.Ctx) error {
	result,err:= eg.service.Get()
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message":err.Error(),
		})
	}

	response := dto.ListModelToResponse(result)
	return f.Status(200).JSON(fiber.Map{
		"message":"success",
		"data":response,
	})
}