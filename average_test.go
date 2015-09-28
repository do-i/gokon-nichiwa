package main

import "testing"

func Average(t *testing.T) {
	nums := []float64{10, 20}
	result := average(nums)
	if result != 15.0 {
		t.Error("expected 15, got ", result)
	}
}