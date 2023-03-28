package util

import "time"

func CurrTime() *time.Time {
	t := time.Now()
	return &t
}
