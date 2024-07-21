package termio

import (
	"strconv"
	"time"
)

func get_time_seconds() string {
	_, _, seconds := time.Now().Clock()
	return strconv.Itoa(seconds)
}

func get_time_minutes() string {
	_, minutes, _ := time.Now().Clock()
	return strconv.Itoa(minutes)
}

func get_time_hours() string {
	hours, _, _ := time.Now().Clock()
	return strconv.Itoa(hours)
}
