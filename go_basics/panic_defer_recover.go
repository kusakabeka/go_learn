package main

import "fmt"

func main() {
	// defer - откладывает выполнение функции в самый конец
	defer handlerPanic()
	fmt.Println("main()")

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	messages[4] = "message 5" // -> panic: runtime error: index out of range [4] with length 4

	fmt.Println(messages)
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("handlerPanic() успешно выполнилась")
}
