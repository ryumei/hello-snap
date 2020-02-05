package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Hello Michael"
	actual := greeting("Michael")
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
func TestGreetingWithoutArg(t *testing.T) {
	expected := "Hello world"
	actual := greeting("")
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
