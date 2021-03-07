package FigureFactory

type rectangle struct {
	side1 float64
	side2 float64
}

func (r rectangle) Area() float64 {
	return r.side1 * r.side2
}

func NewRectangle(side1, side2 float64) Figure {
	return &rectangle {
		side1: side1,
		side2: side2,
	}
}
