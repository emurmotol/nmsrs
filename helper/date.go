package helper

import "time"

func DateForHumans(sec int64) string {
	return time.Unix(sec, 0).Format("January 2, 2006")
}
