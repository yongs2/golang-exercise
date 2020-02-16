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

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
