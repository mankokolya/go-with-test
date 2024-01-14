package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("test 2 + 2", func(t *testing.T) {

		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)

	})

	t.Run("test 3 + 2", func(t *testing.T) {
		sum := Add(3, 2)
		expected := 5
		assertCorrectMessage(t, sum, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}
