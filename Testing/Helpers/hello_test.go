package main

import "testing"

func TestHello(t *testing.T) {
	// This fuunction is a Helper
	assertCorrectMessage := func(t *testing.T, got, expected string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()
		if got != expected {
			t.Errorf("Got %s, expected %s", got, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("K-Sato")
		expected := "Hello, K-Sato"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("saying Hellow World when an empty string is spplied", func(t *testing.T) {
		got := Hello("")
		expected := "Hello, World"
		assertCorrectMessage(t, got, expected)
	})
}
