package helper

import (
	"time"
)

func TimeStampToDate(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04:05")
}
