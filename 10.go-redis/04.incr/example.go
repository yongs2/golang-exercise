package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	now := time.Now()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	client.Ping()
	fmt.Println("Connect.Since=", time.Since(now))

	time.Sleep(2 * time.Second)
	now = time.Now()
	for i := 0; i < 10; i++ {
		now1 := time.Now()
		result, err := client.Incr("counter").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("Incr=", result, "Since=", time.Since(now1))
	}
	fmt.Println("End.Since=", time.Since(now))
	time.Sleep(10 * time.Second)
}
