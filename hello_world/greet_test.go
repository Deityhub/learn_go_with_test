package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greet("Jude")
	want := "Hello, Jude"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	} else {
		t.Logf("got %q want %q", got, want)
	}
}
