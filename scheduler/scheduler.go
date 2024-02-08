package scheduler

import (
	"gocroneg/repository"
	"gocroneg/service"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)



func Scheduler(db *gorm.DB,rdb *redis.Client) error {
	repository := repository.NewRepository(db,rdb)
	service := service.NewService(repository)

	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			3*time.Second,
		),
		gocron.NewTask(
			func() {
				service.Get()
			},
		),
	)

	if err != nil {
		return err
	}

	s.Start()
	
	return nil
}
