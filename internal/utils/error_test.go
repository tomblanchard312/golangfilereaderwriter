package utils

import (
	"errors"
	"testing"
)

func TestCheckError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CheckError did not panic on error")
		}
	}()
	CheckError(errors.New("this is an error"))
}

func TestCheckError_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CheckError panicked unexpectedly: %v", r)
		}
	}()
	CheckError(nil)
}
