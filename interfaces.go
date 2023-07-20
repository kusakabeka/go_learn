package main

/*
	Интерфейсы используют, чтобы описать, каким набором методов должна обладать его реализация!
*/

import (
	"fmt"
	"math"
)

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

type Square struct {
	sideLen float32
}

func (s Square) Area() float32 {
	return s.sideLen * s.sideLen
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (s Square) Perimetr() float32 {
	return s.sideLen * 4
}

func (c Circle) Perimetr() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)

	printInterface(square)
	printInterface(circle)
	printInterface("qwe")
	printInterface(22)
	printInterface(true)

	printInter2("wewqewqeqwewqeqeqwe")
	printInter2(23)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

// пустой интерфейс
func printInterface(i interface{}) {
	// приведение типа
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type", value)
	}
}

func printInter2(i interface{}) {
	// приведение типа
	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
	}
	fmt.Println(len(str))
}
