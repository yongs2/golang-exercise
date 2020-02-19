package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var (
	min = 1  // sec
	max = 10 // sec
)

func main() {
	rand.Seed(time.Now().UnixNano())

	go Monitoring()

	client := GetRedisClient()
	client.Ping()

	now := time.Now()
	SetTimer(client, 100)	// 10 일때에는 정상 동작
	fmt.Println("End.Since=", time.Since(now))

	time.Sleep(time.Duration(max + 20) * time.Second)	// Monitoring 이벤트 종료를 위해 대기 중
	fmt.Println("Done...")
}

func GetRedisClient() *redis.Client {
	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	fmt.Println("REDIS=", redisAddr)

	now := time.Now()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	ret := client.Ping()
	fmt.Println("Ping=", ret, "Connect.Since=", time.Since(now))
	return client
}

// TM- + 일련번호 + 현재시간(시분초) + Expire 시간(초)
func SetTimer(client *redis.Client, nLoop int) {
	var key, value string
	var nExpire int
	var tExpire time.Duration
	var err error

	value = "0"
	for i := 0; i < nLoop; i++ {
		nExpire = rand.Intn(max) + min
		now := time.Now()
		key = fmt.Sprintf("TM-%03d-%s-%03d", i, now.Format("20060102150405"), nExpire)
		tExpire = time.Duration(nExpire) * time.Second
		err = client.Set(key, value, tExpire).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("Set(", key, ",", tExpire, ")", "Since=", time.Since(now))
		time.Sleep(10 * time.Millisecond)
	}
}

// Key expired 이벤트를 수신하여 화면에 표시
func Monitoring() {
	var err error

	client := GetRedisClient()

	now := time.Now()
	//CONFIG SET notify-keyspace-events KEA
	err = client.ConfigSet("notify-keyspace-events", "KEx").Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigSet", "Since=", time.Since(now))

	now = time.Now()
	//PSUBSCRIBE '__key*__:expired'
	pubsub := client.PSubscribe("__key*__:expired")
	_, err = pubsub.Receive()
	if err != nil {
		panic(err)
	}
	fmt.Println("PSubscribe", "Since=", time.Since(now))

	var nIndex, nExpire int
	var timeStr string
	eventTime := time.Now()
	ch := pubsub.Channel()
	for msg := range ch {
		now = time.Now()
		fmt.Sscanf(msg.Payload, "TM-%03d-%14s-%03d", &nIndex, &timeStr, &nExpire)
		eventTime, _ = time.Parse("20060102150405", timeStr)
		eventTime = eventTime.Add(time.Duration(nExpire) * time.Second)
		
		// "Channel=", msg.Channel, 
		fmt.Println(now.Format("2006/01/02 15:04:05"), "Payload=", msg.Payload, 
			"Expire=", eventTime.Format("2006/01/02 15:04:05"), "diff=", now.Sub(eventTime))
	}
}
