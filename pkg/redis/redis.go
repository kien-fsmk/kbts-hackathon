package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type RedisClient struct {
	config *viper.Viper
	client *redis.Client
}

func NewRedisClient(config *viper.Viper) *RedisClient {
	redisHost := config.GetString("redis.host")
	redisPort := config.GetString("redis.port")
	redisPassword := config.GetString("redis.password")
	redisDB := config.GetInt("redis.db")

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	return &RedisClient{config: config, client: rdb}
}

func (r *RedisClient) Set(key string, value interface{}) error {
	return r.client.Set(r.client.Context(), key, value, 0).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.client.Context(), key).Result()
}
