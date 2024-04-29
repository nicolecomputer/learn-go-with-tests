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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected %f got %f", want, got)
		}
	}

	t.Run("30x20 on a rectangle", func(t *testing.T) {
		rect := Rectange{30.0, 20.0}

		checkArea(t, rect, 600.0)
	})

	t.Run("10 Radius Circle", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("12 height 6 base Triangle", func(t *testing.T) {
		triangle := Triangle{Base: 12, Height: 6}

		checkArea(t, triangle, 36.0)
	})
}
