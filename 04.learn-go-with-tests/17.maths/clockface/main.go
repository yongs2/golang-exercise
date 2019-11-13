package main

import (
	"04.learn-go-with-tests/17.maths/svg" // 해당 디렉토리로 이동 후 go install 실행
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Writer(os.Stdout, t)
}
