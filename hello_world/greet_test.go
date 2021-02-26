package main

import "testing"

func TestGreeting(t *testing.T) {
	assertion := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Greet someone (Jude)", func(t *testing.T) {
		got := Greet("Jude")
		want := "Hello, Jude"
		assertion(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Greet("")
		want := "Hello, World"
		assertion(t, got, want)
	})
}
