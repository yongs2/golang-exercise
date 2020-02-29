package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bohov/rate"
)

func main() {
	redisAddr := "redis://" + os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	now := time.Now()
	limiter, err := rate.NewLimiter(redisAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect.Since=", time.Since(now))

	now = time.Now()
	for i := 0; i < 100; i++ {
		// 10 requests in 1 seconds
		res := limiter.Allow("project:123", 10, 1*time.Second)
		fmt.Println(time.Now().Format("15:04:05.999"), "limiter=", res)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("End.Since=", time.Since(now))
}
