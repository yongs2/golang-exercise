package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("YYYY-MM-DD: %s\n", now.Format("2006-01-02"))

	fmt.Printf("YY-MM-DD: %s\n", now.Format("06-01-02"))

	fmt.Printf("YYYY-#{MonthName}-DD: %s\n", now.Format("2006-Jan-02"))

	fmt.Printf("HH:MM:SS: %s\n", now.Format("03:04:05"))

	fmt.Printf("HH:MM:SS Millisecond: %s\n", now.Format("03:04:05 .999"))

	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS: %s\n", now.Format("2006-Jan-02 Monday 03:04:05"))

	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset: %s\n", now.Format("2006-Jan-02 Monday 03:04:05 PM MST -07:00"))
}
