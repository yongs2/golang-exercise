package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/vearne/ratelimit"
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

	now = time.Now()

	// 10 operations per second
	var throughput int = 10
	var alg int = ratelimit.TokenBucketAlg // ratelimit.CounterAlg
	// func NewRedisRateLimiter(client *redis.Client, key string, duration time.Duration, throughput int, batchSize int, alg int) (*RedisRateLimiter, error)
	limiter, _ := ratelimit.NewRedisRateLimiter(client,
		"project:123",
		1*time.Second,
		throughput,
		10,
		alg,
	)
	for i := 0; i < 100; i++ {
		res := limiter.Take()
		fmt.Println(time.Now().Format("15:04:05.999"), "limiter=", res)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("End.Since=", time.Since(now))
}
