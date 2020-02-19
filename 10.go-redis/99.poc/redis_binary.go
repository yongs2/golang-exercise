package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	var key string = "UE0001"

	data, err := ReadBinary("golang-400.jpg")
	if err != nil {
		panic(err)
	}

	fmt.Println("ReadBinary=", len(data))

	client := GetRedisClient()

	now := time.Now()
	for i:=0; i<10; i++ {
		RPushBinary(client, key, data)
	}
	fmt.Println("End.Since=", time.Since(now))

	var nLen int = 0
	var filename string
	for i:=0; i<100; i++ {
		data, err = LPopBinary(client, key)
		if err != nil {
			fmt.Println("LPopBinary.Err=", err.Error())
			break;
		}
		filename = fmt.Sprintf("%s-%03d.jpg", key, i)
		nLen, err = WriteBinary(filename, data)
		if err != nil {
			fmt.Println("WriteBinary.Err=", err.Error())
			break;
		}
		fmt.Println("IDX=", i, "Binary=", len(data), "Write=", nLen)
	}
}

func WriteBinary(filename string, data []byte) (int, error) {
	var nLen int = 0

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0555)
	if err != nil {
		return -1, err
	}

	nLen, err = file.Write(data)
	if err != nil {
		return -1, err
	}

	file.Close()
	return nLen, err
}

func ReadBinary(filename string) ([]byte, error) {
	const maxBytes = 40*1024
	var minSize int = 0

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fmt.Println("stat=", stat.Name(), stat.Size())
	if stat.Size() > maxBytes {
		return nil, err
	} else if stat.Size() == maxBytes {
		minSize = maxBytes
	} else {
		minSize = int(stat.Size())
	}

	data := make([]byte, minSize)
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read %d, Max %d\n", count, maxBytes)

	file.Close()
	return data, nil
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

func RPushBinary(client *redis.Client, key string, data []byte) {
	now := time.Now()
	intCmd := client.RPush(key, data)
	if err := intCmd.Err(); err != nil {
		panic(err)
	}
	fmt.Println("RPush(", key, ")=", intCmd.Val(), ", Since=", time.Since(now))
}

func LPopBinary(client *redis.Client, key string) (data []byte, err error) {
	now := time.Now()
	stringCmd := client.LPop(key)
	if err := stringCmd.Err(); err != nil {
		return nil, err
	}
	data, err = stringCmd.Bytes()
	if err != nil {
		return nil, err
	}
	fmt.Println("LPop(", key, ")=", len(data), ", Since=", time.Since(now))
	return data, err
}
