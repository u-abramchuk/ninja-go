package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	s := Sum(3, 5)

	if s != 8 {
		t.Error("Expected 8, got", s)
	}
}
