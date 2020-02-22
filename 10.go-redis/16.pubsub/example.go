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

	pubsub := client.Subscribe("mychannel1")
	// wait
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}
	ch := pubsub.Channel()

	var msg string
	for i:=0; i<10; i++ {
		msg = fmt.Sprintf("hello-%02d", i)
		err = client.Publish("mychannel1", msg).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("Publish:", msg, "Time:", time.Now().Format("15:04:05.999999"))
	}

	time.AfterFunc(time.Second, func() {
		_ = pubsub.Close()
		fmt.Println("Pubsub.Close()", "Time:", time.Now().Format("15:04:05.999999"))
	})

	for msg := range ch {
		fmt.Println("Channel:", msg.Channel, msg.Payload, "Time:", time.Now().Format("15:04:05.999999"))
	}
	fmt.Println("End.Since=", time.Since(now))
}
