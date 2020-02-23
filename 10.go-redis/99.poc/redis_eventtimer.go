package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-redis/redis"
)

const (
	min = 1  // sec
	max = 10 // sec
)

type EventTimer struct {
	KeyId   string
	SetTime time.Time
	Expire  int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go Monitoring()

	client := GetRedisClient()

	// flush
	client.FlushDB()

	now := time.Now()
	SetTimer(client, 100)
	fmt.Println("End.Since=", time.Since(now))

	time.Sleep(time.Duration(max+20) * time.Second) // Monitoring 이벤트 종료를 위해 대기 중
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
	var key, keyId string
	var nExpire int
	var err error
	var data []byte

	for i := 0; i < nLoop; i++ {
		nExpire = rand.Intn(max) + min
		now := time.Now()
		tExpire := now.Add(time.Second * time.Duration(nExpire))
		keyId = fmt.Sprintf("TM-%03d-%s-%03d", i, now.Format("20060102150405"), nExpire)
		key = fmt.Sprintf("TM-%s", tExpire.Format("20060102150405"))

		event := &EventTimer{
			KeyId:   keyId,
			SetTime: now,
			Expire:  nExpire,
		}
		data, err = json.Marshal(event)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(data))

		intCmd := client.RPush(key, data)
		if err = intCmd.Err(); err != nil {
			panic(err)
		}
		fmt.Println("RPush(", key, ")=", intCmd.Val(), "KeyId=", keyId, ", Since=", time.Since(now))
	}
}

// 매 초 마다 Queue 에서 데이터를 추출하여 화면에 표시
func Monitoring() {
	var err error
	var nWaitMsec int = 100
	var data []byte
	var key string
	var event EventTimer

	client := GetRedisClient()

	for true {
		now := time.Now()
		key = fmt.Sprintf("TM-%s", now.Format("20060102150405"))

		for true {
			stringCmd := client.LPop(key)
			if err = stringCmd.Err(); err == nil {
				data, err = stringCmd.Bytes()
				if err == nil {
					err = json.Unmarshal(data, &event)
					if err == nil {
						fmt.Println(now.Format("2006/01/02/ 15:04:05"), "LPop(", key, ")=", event.KeyId, "Set:", event.SetTime.Format("20060102150405"), ", Since=", time.Since(now))
					}
				}
			} else {
				break
			}
		}
		time.Sleep(time.Duration(nWaitMsec) * time.Millisecond)
	}
}
