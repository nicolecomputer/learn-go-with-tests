package smi

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("10x10 on a rectangle", func(t *testing.T) {
		expected := 40.0

		rect := Rectange{10.0, 10.0}
		got := Perimeter(rect)

		if expected != got {
			t.Errorf("Expected %f got %f", expected, got)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("30x20 on a rectangle", func(t *testing.T) {
		expected := 600.0

		rect := Rectange{30.0, 20.0}
		got := rect.Area()

		if expected != got {
			t.Errorf("Expected %f got %f", expected, got)
		}
	})

	t.Run("10 Radius Circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
