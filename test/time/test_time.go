package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	ts := t.Format("2006-01-02 15:04:05")
	fmt.Println(ts)

	i64 := DateToUnix(ts)
	fmt.Println(i64)

	time := DateToTime(ts)
	fmt.Println(time.Format("2006-01-02 15:04:05"))
}

func DateToUnix(str string) int64 {
	timestamp := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(timestamp, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func DateToTime(str string) time.Time {
	template := "2006-01-02 15:04:05"
	t, err := time.Parse(template, str)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
