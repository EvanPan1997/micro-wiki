package utils

import (
	"log"
	"time"
)

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StringToTime(s string) time.Time {
	template := "2006-01-02 15:04:05"
	t, err := time.Parse(template, s)
	if err != nil {
		log.Println(err)
	}
	return t
}
