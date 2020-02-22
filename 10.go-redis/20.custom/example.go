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

	// flush
	client.FlushDB()

	var v string
	var err error
	var key = "keys_does_not_exist"

	// 0.set
	result := client.Set(key, "40", 0)
	fmt.Printf("Set: %q %s\n", result.Val(), result.Err())

	// 1. define function
	now = time.Now()
	Get := func(rdb *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		rdb.Process(cmd)
		return cmd
	}
	v, err = Get(client, key).Result()
	fmt.Printf("Get: %q %s\n", v, err)
	fmt.Println("End.Since=", time.Since(now))

	// 2. Do cmd
	now = time.Now()
	v, err = client.Do("get", key).Text()
	fmt.Printf("Do: %q %s\n", v, err)
	fmt.Println("End.Since=", time.Since(now))
}
