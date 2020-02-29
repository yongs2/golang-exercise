package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
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

	// flush
	client.FlushDB()

	now = time.Now()
	limiter := redis_rate.NewLimiter(client)
	for i := 0; i < 100; i++ {
		res, err := limiter.AllowN("project:123", redis_rate.PerSecond(10), 1)
		if err != nil {
			panic(err)
		}
		fmt.Println(time.Now().Format("15:04:05.999"), "limiter=", res.Allowed, res.Remaining,
			"ResetAfter:", res.ResetAfter, "RetryAfter:", res.RetryAfter)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("End.Since=", time.Since(now))
}
