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

	client.FlushDB()
	now = time.Now()

	IncrByXX := redis.NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return redis.call("INCRBY", KEYS[1], ARGV[1])
		end
		return false
	`)

	// no key, fail to IncrByXX
	n, err := IncrByXX.Run(client, []string{"xx_counter"}, 2).Result()
	fmt.Println("IncrByXX.Run", "Result=", n, ", Err=", err, ", Time:", time.Now().Format("15:04:05.999999"))

	// set
	err = client.Set("xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	// IncrByXX
	n, err = IncrByXX.Run(client, []string{"xx_counter"}, 2).Result()
	fmt.Println("IncrByXX.Run", "Result=", n, ", Err=", err, ", Time:", time.Now().Format("15:04:05.999999"))

	fmt.Println("End.Since=", time.Since(now))
}
