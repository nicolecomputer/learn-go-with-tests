package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaaa"

		if got != expected {
			t.Errorf("Expected %s got %s", expected, got)
		}
	})

	t.Run("repeat z", func(t *testing.T) {
		got := Repeat("z")
		expected := "zzzzz"

		if got != expected {
			t.Errorf("Expected %s got %s", expected, got)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("z")
	fmt.Println(repeated)
	// Output: zzzzz
}
