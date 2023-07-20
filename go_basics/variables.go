package main

import "fmt"

func main() {

	var message = "123"
	var number float32
	var b bool
	// byte - синоним типа uint8
	// rune - синоним типа int32
	// message := []byte("asd")
	number = 12.2

	var a1 byte = 97
	fmt.Println(a1)
	var a2 rune = 'a'
	fmt.Println(a2)

	fmt.Println(message)
	fmt.Println(number)
	fmt.Println(b)

}
