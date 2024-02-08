package service

import (
	"gocroneg/model"
	"gocroneg/repository"
)

type service struct {
	repository repository.RepositoryInterface
}

type ServiceInterface interface {
	Insert(data model.User) error
	Get() ([]model.User, error)
}

func NewService(repository repository.RepositoryInterface) ServiceInterface {
	return &service{
		repository: repository,
	}
}

func (eg *service) Insert(data model.User) error {
	err := eg.repository.Insert(data)
	if err != nil {
		return err
	}

	return nil
}

func (eg *service) Get() ([]model.User, error) {
	data, err := eg.repository.Get()
	if err != nil {
		return nil, err
	}

	return data, nil
}
