package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2+2", func(t *testing.T) {
		sum := Adder(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("Unexpected, got %d but expected %d", sum, expected)
		}
	})
	t.Run("5+9", func(t *testing.T) {
		sum := Adder(5, 9)
		expected := 14

		if sum != expected {
			t.Errorf("Unexpected, got %d but expected %d", sum, expected)
		}
	})
}

func ExampleAdder() {
	sum := Adder(2, 7)
	fmt.Println(sum)
	// Output: 9
}
