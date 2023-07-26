package main

import "fmt"

func main() {
	a := CreatePtr()
	fmt.Println(*a)

	*a = 100
	fmt.Println(*a)
}

func CreatePtr() *int {
	a := 5
	return &a
}
