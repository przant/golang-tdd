package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrecMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Antonio", "")
		want := "Hello, Antonio"

		assertCorrecMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrecMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrecMessage(t, got, want)
	})
}
