package main

import "testing"

func TestGreetings(t *testing.T) {
	expect := "Hello John"
	actual := greetings("John")

	if expect != actual {
		t.Errorf("%v !== %v\n", expect, actual)
	}
}

func TestMultiple(t *testing.T) {
	expect := 4
	actual := multiple(2, 2)

	if expect != actual {
		t.Errorf("%v is not %v\n", expect, actual)
	}
}
