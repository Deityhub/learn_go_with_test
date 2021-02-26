package main

import "testing"

func TestGreeting(t *testing.T) {
	assertion := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("should greet someone (Jude)", func(t *testing.T) {
		got := Greet("Jude", "")
		want := "Hello, Jude"
		assertion(t, got, want)
	})

	t.Run("should say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Greet("", "")
		want := "Hello, World"
		assertion(t, got, want)
	})

	t.Run("should greet someone in spanish", func(t *testing.T) {
		got := Greet("Joana", "Spanish")
		want := "Hola, Joana"
		assertion(t, got, want)
	})

	t.Run("Should greet someone in French", func(t *testing.T) {
		got := Greet("Marcel", "French")
		want := "Bonjour, Marcel"
		assertion(t, got, want)
	})
}
