package test

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"github.com/sdttttt/go-sso/model"
)

func TestExampleNewClient(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "121.40.48.194:10086",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func TestExampleNewClient2(t *testing.T) {
	client := model.RedisConnectInstance()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
