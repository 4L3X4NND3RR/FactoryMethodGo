package FigureFactory

type square struct {
	side1 float64
}

func (s square) Area() float64 {
	return s.side1 * s.side1
}

func NewSquare(side1 float64) Figure {
	return &square{
		side1: side1,
	}
}
