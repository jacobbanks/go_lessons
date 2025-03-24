package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to the perople", func(t *testing.T) {
		got := Hello("Banks")
		want := "Hello, Banks"
		assertCorrectMessage(t, got, want)
	})
	t.Run("default to 'Hello, World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	
}
