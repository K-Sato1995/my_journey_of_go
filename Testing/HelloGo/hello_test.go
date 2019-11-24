package main

import "testing"

func TestHelloGo(t *testing.T) {
	got := Hello()
	expected := "Hello World"

	if got != expected {
		t.Errorf("Got %s, expected %s", got, expected)
	}
}
