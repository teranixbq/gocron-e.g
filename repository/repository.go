package repository

import (
	"gocroneg/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	Insert(data model.User) error
	Get() (model.User, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (eg *repository) Insert(data model.User) error {

	tx := eg.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (eg *repository) Get() (model.User, error) {
	dataUser := model.User{}

	tx := eg.db.Find(&dataUser)
	if tx.Error != nil {
		return model.User{},tx.Error
	}

	return dataUser,nil
}

