package util

import (
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	{
		var i = 3
		Retry(3, time.Second, func() error {
			i--
			return errors.New("")
		})

		if i != 0 {
			t.Error("faile")
		}
	}

	{
		var i = 3
		Retry(3, time.Second, func() error {
			i--
			return RetryStop{errors.New("err")}
		})

		if i != 2 {
			t.Error("failed")
		}
	}

	{
		var i = 3
		Retry(3, time.Second, func() error {
			i--
			return nil
		})

		if i != 2 {
			t.Error("faile")
		}
	}
}
