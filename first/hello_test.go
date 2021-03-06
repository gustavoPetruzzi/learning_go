package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("crack", "English")
		want:= "Hello crack"
		assertCorrectMessage(t, got, want);

	})

	t.Run("Saying hello to people in spanish", func(t *testing.T) {
		got:= Hello("crack", "Spanish")
		want:= "Hola crack"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in french", func(t *testing.T) {
		got:= Hello("crack", "French")
		want:= "Bonjour crack"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want:= "Hello world"
		assertCorrectMessage(t, got, want)
	})

}