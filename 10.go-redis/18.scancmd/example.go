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

	// create key*
	client.FlushDB()
	now = time.Now()
	for i := 0; i < 33; i++ {
		err := client.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), 0).Err()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Set.Since=", time.Since(now))

	// scan
	now = time.Now()
	var idx = 0
	iter := client.Scan(0, "", 0).Iterator()
	for iter.Next() {
		fmt.Println(iter.Val())
		idx = idx + 1
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	fmt.Println("IDX=", idx, "End.Since=", time.Since(now))
}
