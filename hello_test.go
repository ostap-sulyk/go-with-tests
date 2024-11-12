package main

import "testing"

func TestHello(t *testing.T) {
	name := "Chris"
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello(name)
		want := englishHelloPrefix + name

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertCorrectMessege(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
