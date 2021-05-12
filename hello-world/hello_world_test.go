package main

import "testing"

func TestHelloWorld(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, result string, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("got %q want %q", result, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		result := HelloWorld("Rhuan")
		expected := "Hello, Rhuan"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		result := HelloWorld("")
		expected := "Hello, World"

		assertCorrectMessage(t, result, expected)
	})
}
