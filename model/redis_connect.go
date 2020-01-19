package model

import (
	"github.com/go-redis/redis"
	"sync"
)

var client *redis.Client
var once sync.Once

func RedisConnectInstance() *redis.Client {
	once.Do(
		func() {
			client = redis.NewClient(&redis.Options{
				Addr:     "121.40.48.194:10086",
				Password: "", // no password set
				DB:       0,  // use default DB
			})
		})

	return client
}
