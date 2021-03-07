package FigureFactory

import "math"

type triangle struct {
	side1 float64
	side2 float64
	side3 float64
}

func (t triangle) Area() float64 {
	P := (t.side1 + t.side2 + t.side3) / 2
	return math.Sqrt(P * (P - t.side1) * (P - t.side2) * (P - t.side3))
}

func NewTriangle(side1, side2, side3 float64) Figure {
	return &triangle{
		side1: side1,
		side2: side2,
		side3: side3,
	}
}
