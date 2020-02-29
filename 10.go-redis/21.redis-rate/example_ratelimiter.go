package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	ratelimiter "github.com/teambition/ratelimiter-go"
)

type redisClient struct {
	*redis.Client
}

func (c *redisClient) RateDel(key string) error {
	return c.Del(key).Err()
}

func (c *redisClient) RateEvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return c.EvalSha(sha1, keys, args...).Result()
}

func (c *redisClient) RateScriptLoad(script string) (string, error) {
	return c.ScriptLoad(script).Result()
}

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
	limiter := ratelimiter.New(ratelimiter.Options{
		Max:      10,
		Duration: time.Second,
		Client:   &redisClient{client},
	})

	for i := 0; i < 100; i++ {
		res, err := limiter.Get("project:123")
		if err != nil {
			panic(err)
		}
		fmt.Println(time.Now().Format("15:04:05.999"), "limiter=", res.Total, res.Remaining,
			"Duration:", res.Duration, "Reset:", res.Reset)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("End.Since=", time.Since(now))
}
