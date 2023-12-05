package utils

import "time"

func GetCurrentTime() time.Time {
	current := time.Now().UTC()

	location := time.FixedZone("Asia/Jakarta", 7*60*60)

	return current.In(location)
}
