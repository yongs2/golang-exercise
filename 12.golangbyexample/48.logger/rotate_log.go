package main

import (
	"fmt"
	"log"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	//"github.com/sirupsen/logrus"
)

func initiLogger() {
	path := "/var/log/service/old.UTC."
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
		rotatelogs.WithLinkName("/var/log/service/current"),
		rotatelogs.WithMaxAge(time.Second*10),
		rotatelogs.WithRotationTime(time.Second*1),
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	log.SetOutput(writer)
	return
}

func main() {
	initiLogger()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 1)
		log.Printf("Hello, World!")
	}
	fmt.Scanln()
}
