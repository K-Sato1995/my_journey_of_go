package main

import "testing"

func TestCheckScore(t *testing.T) {
	expect := "A"
	actual := checkScore(90)

	if expect != actual {
		t.Errorf("Something ain't right")
	}
}
