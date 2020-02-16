package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// Author ...
type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Go Redis Tutorial")

	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	// ping, pong
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// set json
	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Setting, Err=", err)

	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Getting, val=", val)
}
