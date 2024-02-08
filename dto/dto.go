package dto

import "gocroneg/model"

type Request struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
	Email   string `json:"email"`
}

type Response struct {
	Id	  uint   `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
	Email   string `json:"email"`
}

func ModelToResponse(data model.User) Response {
	return Response{
		Id:      data.ID,
		Name:    data.Name,
		Age:     data.Age,
		Address: data.Address,
		Telp:    data.Telp,
		Email:   data.Email,
	}
}

func ListModelToResponse (data []model.User) []Response{
	var result []Response 
	for _, v := range data {
		result = append(result, ModelToResponse(v))
	}
	return result
}

func RequesToModel(data Request) model.User {
	return model.User{
		Name:    data.Name,
		Age:     data.Age,
		Address: data.Address,
		Telp:    data.Telp,
		Email:   data.Email,
	}
}

