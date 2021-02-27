package iteration

import "testing"

func TestIterate(t *testing.T) {
	t.Run("should print hello five times", func(t *testing.T) {
		got := Iterate("hello", 5)
		want := "hellohellohellohellohello"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("should print hello once, when zero value is passed to it", func(t *testing.T) {
		got := Iterate("hello", 0)
		want := "hello"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("should print hello once, when negative value is passed to it", func(t *testing.T) {
		got := Iterate("hello", -2)
		want := "hello"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 5)
	}
}
