package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func InitRedis()*redis.Client {
	godotenv.Load()
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("RADDRS"),
		Username: os.Getenv("RUSER"),
		Password: os.Getenv("RPASS"),
	})
	return client
}