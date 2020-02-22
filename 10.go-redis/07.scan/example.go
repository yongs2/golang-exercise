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
	for i := 0; i < 33; i++ {
		err := client.Set(fmt.Sprintf("key%d", i), "value", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Set.Since=", time.Since(now))

	now = time.Now()
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "key*", 10).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		fmt.Println("keys=", keys, "cursor=", cursor, "n=", n)
		if cursor == 0 {
			break
		}
	}
	fmt.Println("Scan.Since=", time.Since(now))
	fmt.Printf("found %d keys\n", n)
	time.Sleep(10 * time.Second)
}
