package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"time"
)

var c *redis.Client

func init() {
	c = redis.NewClient(&redis.Options{
		Addr:         "47.99.3.33:16379",
		Password:     "laystbzqzygwcs",
		DB:           0,
		MaxRetries:   3,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	if _, err := c.Ping().Result(); err != nil {
		panic(errors.Wrap(err, "Could not connect to redis db0"))
	}
}

func Incr(key string) bool {
	keyStr := fmt.Sprintf("go_id_gen_%s", key)
	_, err := c.Incr(keyStr).Result()
	return err == nil
}
