package main

import (
	"errors"
	"fmt"
)

var msg string

/*
срабатывает при инициализации пакета
используется для инициализации
вызывается до main
*/
func init() {
	msg = "from init()"
}

func main() {
	fmt.Println(msg)
	message := "text"
	chngeMsg(&message)
	fmt.Println(message)

	var p *int

	number := 5
	p = &number
	*p = 10

	fmt.Println(number)

	arr := [3]string{"1", "2", "3"}

	printMessages(arr)
	fmt.Println(arr)

	//fmt.Println(arr[4]) - invalid array index

	// SLICES
	// хранит в себе ссылку на массив
	sl := []string{"a", "b", "c"}

	printMessagesSlice(sl)
	fmt.Println(sl)

	/*	PANIC
		var sl1 []string
		sl1[0] = "text"
		fmt.Println(sl1)*/
	// slice with MAKE
	sl1 := make([]string, 5)

	//sl1[6] = "6" - index out of range

	sl1 = append(sl1, "6")
	sl1 = append(sl1, "7")
	sl1 = append(sl1, "8")
	sl1 = append(sl1, "9")
	sl1 = append(sl1, "10")

	fmt.Println(sl1)

	fmt.Println("len:", len(sl1))
	fmt.Println("cap:", cap(sl1))
	sl1[0] = "1"
	fmt.Println(sl1)

	// двумерные
	matrix := make([][]int, 10)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][x] = x
		}
		fmt.Println(matrix[x])
	}
}

// здесь передаем ссылку на sl
func printMessagesSlice(message []string) error {
	if len(message) == 0 {
		return errors.New("empty array")
	}

	message[1] = "d"

	fmt.Println(message)

	return nil
}

func printMessages(message [3]string) error {
	if len(message) == 0 {
		return errors.New("empty array")
	}

	message[1] = "5"

	fmt.Println(message)

	return nil
}

func chngeMsg(message *string) {
	*message += " (from printMsg)"
	fmt.Println(*message)
}
