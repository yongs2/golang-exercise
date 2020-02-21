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

	now = time.Now()

	conn := client.Conn()

	err := conn.ClientSetName("foobar").Err()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		go client.Ping()
	}

	s, err := conn.ClientGetName().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ClientGetName=", s, "err=", err, "Since=", time.Since(now))

	time.Sleep(10 * time.Second)
}
