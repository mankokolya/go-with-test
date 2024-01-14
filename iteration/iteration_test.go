package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat a 5 times", func(t *testing.T) {
		repeated := RepeatTimes("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat a 3 times", func(t *testing.T) {
		repeated := RepeatTimes("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatTimes("a", 5)
	}
}
