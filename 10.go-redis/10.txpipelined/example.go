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
	var incr *redis.IntCmd
	_, err := client.TxPipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("tx_pipelined_counter")
		pipe.Expire("tx_pipelined_counter", time.Minute)
		return nil
	})
	fmt.Println("Val=", incr.Val(), "err=", err, "Pipe.Since=", time.Since(now))

	time.Sleep(10 * time.Second)
}
