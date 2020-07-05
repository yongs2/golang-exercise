package main

import (
	"log"
	"regexp"
)

var input string
var r = regexp.MustCompile("(?P<ISO8601>(?P<year>\\d{4})(\\-W((?P<week>\\d{1,2})\\-(?P<weekday>\\d)?)|\\-(?P<month>\\d{2})\\-(?P<day>\\d{2})(T(?P<hour>\\d{2}):(?P<min>\\d{2})(:(?P<sec>\\d{2})(\\+\\d{2}:\\d{2}Z?)?)?)?|\\-(?P<yearday>\\d{1,3})))")

func main() {
	input = `
2014-06-06
2014-11-01
2014-11-01T10:01:13+00:00
2014-11-01T10:01:13Z
2014-W44
2014-W44-6
2014-305
2014-23
`

	match := r.FindAllString(input, -1)
	for i := 0; i < len(match); i++ {
		submatch := r.FindStringSubmatch(match[i])
		log.Printf("\nInput: %#v\n", match[i])
		for i, name := range r.SubexpNames() {
			if name != "" && submatch[i] != "" {
				log.Printf("\t%s\t %s\n", name, submatch[i])
			}
		}
	}
}
