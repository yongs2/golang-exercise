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
	fmt.Println("Connect.Since=", time.Since(now))

	case01(client, 10)
	case02(client, 10)
}

func case01(client *redis.Client, counter int) {
	now := time.Now()
	for i := 0; i < counter; i++ {
		now1 := time.Now()
		pipe := client.Pipeline()
		incr := pipe.Incr("pipeline_counter")
		pipe.Expire("pipeline_counter", time.Hour)
		_, err := pipe.Exec()
		fmt.Println("Incr=", incr.Val(), "Err=", err, "Since=", time.Since(now1))
	}

	fmt.Println("End.Since=", time.Since(now))
}

func case02(client *redis.Client, counter int) {
	now := time.Now()

	pipe := client.Pipeline()
	var incr *redis.IntCmd
	for i := 0; i < counter; i++ {
		now1 := time.Now()

		incr = pipe.Incr("pipeline_counter")
		pipe.Expire("pipeline_counter", time.Hour)
		fmt.Println("Incr=", incr.Val(), "Since=", time.Since(now1))
	}
	_, err := pipe.Exec()
	fmt.Println("Incr=", incr.Val(), "Err=", err, "End.Since=", time.Since(now))
}
