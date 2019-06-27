package main

import "testing"

func TestCalAverage(t *testing.T) {
	expect := 5
	actual := calAve([]int{4, 5, 6})

	if expect != actual {
		t.Errorf("Something ain't right")
	}
}
