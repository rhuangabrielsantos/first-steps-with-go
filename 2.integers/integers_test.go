package integers

import "testing"

func TestOperationsWithIntegers(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, result, expected int) {
		t.Helper()

		if result != expected {
			t.Errorf("expected '%d', result '%d'", expected, result)
		}
	}

	t.Run("given two numbers, when sum function is called, should sum numbers", func(t *testing.T) {
		result := Sum(2, 2)
		expected := 4

		assertCorrectMessage(t, result, expected)
	})

	t.Run("given two numbers, when subtration function is called, should subtration numbers", func(t *testing.T) {
		result := Subtraction(5, 2)
		expected := 3

		assertCorrectMessage(t, result, expected)
	})
}
