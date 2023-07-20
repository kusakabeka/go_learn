package shape

/*
	Интерфейсы используют, чтобы описать, каким набором методов должна обладать его реализация!
*/

import "math"

type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimetr() float32
}

// создаем Квадрат
type Square struct {
	sideLen float32
}

func (s Square) Area() float32 {
	return s.sideLen * s.sideLen
}

func (s Square) Perimetr() float32 {
	return 4 * s.sideLen
}

func NewSquare(length float32) Square {
	return Square{
		sideLen: length,
	}
}

// создаем Круг
type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimetr() float32 {
	return 2 * math.Pi * c.radius
}
