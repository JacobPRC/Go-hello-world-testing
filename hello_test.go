package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func TestGoodbye(t *testing.T) {

	errorMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying goodbye to people", func(t *testing.T) {
		got := Goodbye("Chris", "")
		want := "Goodbye, Chris"
		errorMessage(t, got, want)
	})

	t.Run("Saying goodbye to life", func(t *testing.T) {
		got := Goodbye("", "")
		want := "Goodbye, Cruel World"
		errorMessage(t, got, want)
	})

	t.Run("goodbye in Spanish", func(t *testing.T) {
		got := Goodbye("Elodie", "Spanish")
		want := "Adios, Elodie"
		errorMessage(t, got, want)
	})
}
