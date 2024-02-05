package dto

import "gocroneg/model"

type Request struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
	Email   string `json:"email"`
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
