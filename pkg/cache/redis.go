package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Redis interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
}

type RedisCache struct {
	Addr    string
	Db      int
	Expires time.Duration
}

func NewRedisCache(addr string, db int, expires time.Duration) Redis {
	return &RedisCache{
		Addr:    addr,
		Db:      db,
		Expires: expires,
	}
}

func (r *RedisCache) Client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: "",
		DB:       r.Db,
	})
}

func (r *RedisCache) Get(key string) (string, error) {
	client := r.Client()
	res, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (r *RedisCache) Set(key string, value interface{}) error {
	client := r.Client()

	client.Set(key, value, r.Expires)
	return nil
}
