package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repet a 5 times", func(t *testing.T) {
		assertCorrectMessege(t, Repeat("a", 5), "aaaaa")
	})

	t.Run("Repeat a 0 times", func(t *testing.T) {
		assertCorrectMessege(t, Repeat("a", 0), "")
	})

	t.Run("Do not repeat when have negative number", func(t *testing.T) {
		assertCorrectMessege(t, Repeat("a", -1), "")
	})

}
func assertCorrectMessege(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q but want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
