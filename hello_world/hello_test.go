package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to the perople", func(t *testing.T) {
		got := Hello("Banks", "english")
		want := "Hello, Banks"
		assertCorrectMessage(t, got, want)
	})
	t.Run("default to 'Hello, World'", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("return output in Spanish", func(t *testing.T) {
		got := Hello("Banks", "spanish")
		want := "Hola, Banks"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	
}
