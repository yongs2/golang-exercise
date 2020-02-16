package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		fmt.Println("Setting, Err=", err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		fmt.Println("Getting, Err=", err)
	}
	fmt.Println("key=", val)

	val2, err := client.Get("missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
}
