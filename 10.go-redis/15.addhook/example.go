package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type redisHook struct{}

var _ redis.Hook = redisHook{}

func (redisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	fmt.Printf("starting processing: <%s>\n", cmd)
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	fmt.Printf("finished processing: <%s>\n", cmd)
	return nil
}

func (redisHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	fmt.Printf("pipeline starting processing: %v\n", cmds)
	return ctx, nil
}

func (redisHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	fmt.Printf("pipeline finished processing: %v\n", cmds)
	return nil
}

func main() {
	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	now := time.Now()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	client.AddHook(redisHook{})
	fmt.Println("Connect.Since=", time.Since(now))

	now = time.Now()
	_, err := client.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Ping()
		pipe.Ping()
		return nil
	})
	fmt.Println("Pipelined.Error=", err, "End.Since=", time.Since(now))
}
