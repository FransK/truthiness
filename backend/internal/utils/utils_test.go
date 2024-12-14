package utils

import (
	"math"
	"testing"
)

func TestGetFloat(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panicked: %v", r)
		}
	}()

	f, _ := GetFloat(nil)

	if !math.IsNaN(f) {
		t.Errorf("want math.NaN; got %v", f)
	}
}
