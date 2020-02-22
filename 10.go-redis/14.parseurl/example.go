package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	redisUrl := os.Getenv("REDIS_URL")
	fmt.Println("REDIS_URL=", redisUrl)

	now := time.Now()
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("addr is", opt.Addr)
	fmt.Println("db is", opt.DB)
	fmt.Println("password is", opt.Password)

	client := redis.NewClient(opt)

	status := client.Ping()
	if err := status.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Ping=", status.Val(), "Connect.Since=", time.Since(now))

	time.Sleep(10 * time.Second)
}
