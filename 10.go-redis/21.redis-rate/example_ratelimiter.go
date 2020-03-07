package main

import (
	"fmt"
	"os"
	"time"
	"log"
	"sync"

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
	var wg sync.WaitGroup

	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	log.Println("REDIS=", redisAddr)

	now := time.Now()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	client.Ping()
	log.Println("Connect.Since=", time.Since(now))

	now = time.Now()
	limiter := ratelimiter.New(ratelimiter.Options{
		Max:      10,
		Duration: time.Second,
		Client:   &redisClient{client},
	})

	keys := []string{"key001", "key002"}
	for _, key := range keys {
		wg.Add(1)
		go func(keyVal string) {
			defer wg.Done()
			
			for i := 0; i < 100; i++ {
				res, err := limiter.Get(keyVal)
				if err != nil {
					panic(err)
				}
				fmt.Println(time.Now().Format("15:04:05.999"), keyVal, "limiter=", res.Total, res.Remaining,
					"Duration:", res.Duration, "Reset:", res.Reset)
				time.Sleep(10 * time.Millisecond)
			}
			log.Println("End.Since=", time.Since(now))
		}(key)
	}

	log.Println("Wait....", time.Since(now))
	wg.Wait()
}
