package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("First Test", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d for numbers %v", got, want, numbers)
		}
	})
	t.Run("30 60 90 60 30", func(t *testing.T) {
		numbers := [5]int{30, 60, 90, 60, 30}

		got := Sum(numbers)
		want := 270

		if got != want {
			t.Errorf("got %d want %d for numbers %v", got, want, numbers)
		}
	})
}
