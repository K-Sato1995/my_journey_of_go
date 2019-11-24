package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("K-Sato")
		expected := "Hello, K-Sato"

		if got != expected {
			t.Errorf("Got %s, expected %s", got, expected)
		}
	})

	t.Run("saying Hellow World when an empty string is spplied", func(t *testing.T) {
		got := Hello("")
		expected := "Hello, World"

		if got != expected {
			t.Errorf("Got %s, expected %s", got, expected)
		}
	})
}
