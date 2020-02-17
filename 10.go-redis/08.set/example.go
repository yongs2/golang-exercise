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

	// wait
	time.Sleep(2 * time.Second)

	client.FlushDB()
	// set key, no expiration time
	now = time.Now()
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set.Since=", time.Since(now))

	// key2 will expire in 1 minute
	now = time.Now()
	err = client.Set("key2", "value", time.Minute).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set.Since=", time.Since(now))

	time.Sleep(10 * time.Second)
}
