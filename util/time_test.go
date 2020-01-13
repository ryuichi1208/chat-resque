package util

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	var time, callGetCurrentTime = time.Now().Format("2006-01-02 15:04:05"), GetCurrentTime()
	if time == callGetCurrentTime {
		return
	}
	t.Error("failed")
}
