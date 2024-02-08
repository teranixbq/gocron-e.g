package repository

import (
	"context"
	"encoding/json"
	"gocroneg/model"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

type RepositoryInterface interface {
	Insert(data model.User) error
	Get() ([]model.User, error)
}

func NewRepository(db *gorm.DB, rdb *redis.Client) RepositoryInterface {
	return &repository{
		db:  db,
		rdb: rdb,
	}
}

var (
	ctx = context.Background()
)

func (eg *repository) Insert(data model.User) error {

	tx := eg.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (eg *repository) Get() ([]model.User, error) {

	dataUser := []model.User{}

	val, err := eg.rdb.Get(ctx, "user").Result()
	if err != nil {
		tx := eg.db.Find(&dataUser)
		if tx.Error != nil {
			return nil, tx.Error
		}

		bytes, err := json.Marshal(dataUser)
		if err != nil {
			return nil, err
		}

		
		txr := eg.rdb.SetEx(ctx, "user", bytes, time.Second*5)
		if txr.Err() != nil {
			return nil, txr.Err()
		}

		log.Println("data from db")

		return dataUser, nil
	}

	data := &dataUser
	err = json.Unmarshal([]byte(val), data)

	if err != nil {
		return nil, err
	}

	log.Println("data from redis")
	return *data, nil
}
