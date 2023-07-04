package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Scan(&name, &age)
	message := sayHello(name, age)

	printMessage(message)

	// 	message1, _ := enterTheClub(19) -> _ будет проигнорированно
	message1, entered := enterTheClub(19)
	fmt.Println(message1, entered)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hi, %s! You are %d years old!", name, age)
}

func enterTheClub(age byte) (string, bool) {
	if age >= 18 {
		return "Welcome, but carefully", true
	} else {
		return "You are not 18 years old!", false
	}
}
