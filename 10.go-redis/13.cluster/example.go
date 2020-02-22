package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	redisCluster1 := os.Getenv("REDIS_CLUSTER1")
	redisCluster2 := os.Getenv("REDIS_CLUSTER2")
	redisCluster3 := os.Getenv("REDIS_CLUSTER3")
	fmt.Println("REDIS_CLUSTER=", redisCluster1, redisCluster2, redisCluster3)

	now := time.Now()
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redisCluster1, redisCluster2, redisCluster3},
		Password: "",
	})
	status := client.Ping()
	if err := status.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Connect.Since=", time.Since(now))

	time.Sleep(2 * time.Second)
	now = time.Now()
	for i := 0; i < 100; i++ {
		now1 := time.Now()
		result, err := client.Incr("counter").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("Incr=", result, "Since=", time.Since(now1))
	}
	fmt.Println("End.Since=", time.Since(now))

	// incr, set, etc..
	time.Sleep(10 * time.Second)
}
