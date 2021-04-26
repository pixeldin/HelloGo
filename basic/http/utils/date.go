package utils

import "time"

func GetPresentFormat() string {
	time.LoadLocation("local")
	return time.Now().Format("2006-01-02 15:04:05.99")
}
