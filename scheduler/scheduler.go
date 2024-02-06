package scheduler

import (
	"gocroneg/repository"
	"gocroneg/service"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)



func Scheduler(db *gorm.DB) error {
	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			10*time.Second,
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

	select {}
	
}
