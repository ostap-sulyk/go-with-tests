package helloworld

import "testing"

func TestHello(t *testing.T) {
	name := "Chris"
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello(name, "")
		want := englishHelloPrefix + name

		assertCorrectMessege(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessege(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		name := "Elodie"
		got := Hello(name, spanish)
		want := spanishHelloPrefix + name

		assertCorrectMessege(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("John", french)
		want := franchHelloPrefix + "John"

		assertCorrectMessege(t, got, want)
	})
}

func assertCorrectMessege(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
