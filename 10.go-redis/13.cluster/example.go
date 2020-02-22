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
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"172.17.0.3:6379", "172.17.0.4:6380", "172.17.0.5:6381"},
		Password: "",
	})
	client.Ping()
	fmt.Println("Connect.Since=", time.Since(now))

	// incr, set, etc..
	time.Sleep(10 * time.Second)
}
