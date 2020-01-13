package util

import (
	"errors"
	"testing"
)

func TestCheckError(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	CheckError(errors.New("err"))
	t.Error("failed")
}
