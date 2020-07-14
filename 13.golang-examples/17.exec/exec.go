package main

import (
	"log"
	"os/exec"
	"strings"
)

func pipe_cmd(cmds ...*exec.Cmd) ([]byte, error) {
	for i, cmd := range cmds[:len(cmds)-1] {
		out, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		cmd.Start()
		cmds[i+1].Stdin = out
	}

	ret, err := cmds[len(cmds)-1].Output()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func main() {
	datestr, err := exec.Command("date").Output()
	if err != nil {
		log.Println(err)
	}
	log.Print("Date String: ", string(datestr))

	var datecmd *exec.Cmd
	var lscmd *exec.Cmd

	datecmd = exec.Command("date", "+%d.%m.%Y %H:%M:%S")
	grep1 := exec.Command("grep", "-o", "\\d\\d.\\d\\d.\\d\\d\\d\\d")
	date, err1 := pipe_cmd(datecmd, grep1)

	datecmd = exec.Command("date", "+%d.%m.%Y %H:%M:%S")
	grep2 := exec.Command("grep", "-o", "\\d\\d:\\d\\d:\\d\\d")
	time, err2 := pipe_cmd(datecmd, grep2)

	lscmd = exec.Command("ls", "-ltraR", "./")
	grep3 := exec.Command("grep", "-o", "[^ \\./]\\+\\.[^ \\./]\\+")
	files, err3 := pipe_cmd(lscmd, grep3)

	if err1 != nil {
		log.Print(err1)
	} else {
		log.Print("Date: ", string(date))
	}

	if err2 != nil {
		log.Print(err2)
	} else {
		log.Print("Time: ", string(time))
	}

	counter := make(map[string]int)
	var ok bool
	if err3 != nil {
		log.Print(err3)
	} else {
		strarray := strings.Split(string(files), "\n")
		for i := 0; i < len(strarray); i++ {
			if strings.Trim(strarray[i], " .") != "" {
				if _, ok = counter[strarray[i]]; ok {
					counter[strarray[i]]++
				} else {
					counter[strarray[i]] = 1
				}
			}
		}
		log.Println("Files: ", counter)
	}
}
