package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")
	
	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		Password: "",
		DB: 0,
	})

	// ping, pong
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// adding values to redis
	err = client.Set("name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Setting, Err=", err)

	// getting values
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Getting, val=", val)
}
