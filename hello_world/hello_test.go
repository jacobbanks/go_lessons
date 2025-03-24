package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Banks")
	want := "Hello, Banks"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
