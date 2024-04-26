package smi

import "math"

type Rectange struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectange) float64 {
	return 2*rect.Width + 2*rect.Height
}

func (rect Rectange) Area() float64 {
	return rect.Width * rect.Height
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
