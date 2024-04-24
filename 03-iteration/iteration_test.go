package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat A", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaaa"

		if got != expected {
			t.Errorf("Expected %s got %s", expected, got)
		}
	})
}
