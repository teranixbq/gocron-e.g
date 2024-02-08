package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Age     string
	Address string
	Telp    string
	Email   string
}

func (User) TableName() string {
	return "user"
}