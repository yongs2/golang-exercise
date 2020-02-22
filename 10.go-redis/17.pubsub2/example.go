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

	now = time.Now()

	pubsub := client.Subscribe("mychannel2")
	defer pubsub.Close()
	fmt.Println("client.Subscribed", "Time:", time.Now().Format("15:04:05.999999"))

	for i := 0; i < 3; i++ {
		msgi, err := pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			fmt.Println("ReceiveTimeout", "Time:", time.Now().Format("15:04:05.999999"))
			break
		}

		switch msg := msgi.(type) {
		case *redis.Subscription:
			fmt.Println("subscribed to", msg.Channel, "Time:", time.Now().Format("15:04:05.999999"))

			_, err := client.Publish("mychannel2", "hello").Result()
			if err != nil {
				panic(err)
			}
		case *redis.Message:
			fmt.Println("received", msg.Payload, "from", msg.Channel, "Time:", time.Now().Format("15:04:05.999999"))
		default:
			panic("unreached")
		}
	}

	fmt.Println("End.Since=", time.Since(now))
}
