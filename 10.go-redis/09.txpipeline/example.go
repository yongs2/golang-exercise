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
	// set key, no expiration time
	now = time.Now()
	pipe := client.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Minute)

	/*
		MULTI
		INCR tx_pipeline_counter
		EXPIRE tx_pipeline_counter 60
		EXEC
	*/
	_, err := pipe.Exec()
	fmt.Println("Val=", incr.Val(), "err=", err, "Pipe.Since=", time.Since(now))

	time.Sleep(10 * time.Second)
}
