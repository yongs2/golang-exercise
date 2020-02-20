package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
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

	const routineCount = 100
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			n, err := tx.Get(key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			n++

			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				status := pipe.Set(key, n, 0)
				fmt.Println("==> Set(", key, ",", n, ")=", status.Err())
				return nil
			})
			return err
		}

		for retries := routineCount; retries > 0; retries-- {
			err := client.Watch(txf, key)
			if err != redis.TxFailedErr {
				return err
			}
		}
		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := client.Get("counter3").Int()
	fmt.Println("Val=", n, "err=", err, "Pipe.Since=", time.Since(now))

	time.Sleep(10 * time.Second)
}
