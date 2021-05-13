package interation

import "testing"

func TestInteraction(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("expected '%s', result '%s'", expected, result)
		}
	}

	t.Run("given one character, with repetitions number set to five, should return five times", func(t *testing.T) {
		result := Repeater("a", 6)
		expected := "aaaaaa"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("given one character, with repetitions number set to fifteen, should return fifteen times", func(t *testing.T) {
		result := Repeater("a", 15)
		expected := "aaaaaaaaaaaaaaa"

		assertCorrectMessage(t, result, expected)
	})

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 5)
	}
}
