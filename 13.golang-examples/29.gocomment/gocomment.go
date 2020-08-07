package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

// go build gocomment.go
// ./gocomment go-app.go
func main() {
	if len(os.Args) >= 2 {
		var output string
		var i int

		dateOut, err := exec.Command("/usr/local/go/bin/go build /go/src/01.hello/hello.go").Output()
		if err != nil {
			panic(err)
		}
		fmt.Println("> date:", string(dateOut))

		input := os.Args[1]
		gosrcs, _ := ioutil.ReadFile(input)
		gosrc := string(gosrcs)
		gorun, err := exec.Command("go", "build", input).Output()
		fmt.Println("go build", input, ".Output()=[", gorun, "], err=[", err, "]")
		gostr := string(gorun)
		fmt.Println("gostr=[", gostr, "]")
		aline := regexp.MustCompile("[^\\n]*\\n")
		line, _ := regexp.Compile(input + ":(\\d+): ")
		stre := line.FindAllStringSubmatch(gostr, -1)
		stra := aline.FindAllString(gosrc, -1)
		j := 0
		fmt.Println("len(stre)=", len(stre), "len(stra)=", len(stra))
		for i = 0; i < len(stra); i++ {
			if len(stre) > j {
				lc, _ := strconv.ParseInt(stre[j][1], 10, 0)
				if i == int(lc)-1 {
					output += "//" + stra[i]
					j++
				} else {
					fmt.Println(i)
					fmt.Println(stre[j][1])
					output += stra[i]
				}
			} else {
				output += stra[i]
			}
		}
		outputb := []byte(output)
		ioutil.WriteFile(input, outputb, 0755)
	}
}
