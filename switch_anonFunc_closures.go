package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := prediction("чт")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	fmt.Println(findMinOrMax("max", 1, 22, 34, 51, 25))

	func() {
		fmt.Println("anon func")
	}()
	// closures
	inc := increment()                      // inc - func(), не int
	fmt.Println(inc(), inc(), inc(), inc()) // 1 2 3 4

	fmt.Println(increment2(), increment2(), increment2(), increment2()) // 1 1 1 1
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "01", nil
	case "вт":
		return "02", nil
	case "ср":
		return "03", nil
	case "чт":
		return "04", nil
	default:
		return "невалидный день недели", errors.New("invalid day of the week")
	}
}
func findMinOrMax(word string, numbers ...int) int {

	min := numbers[0]
	max := numbers[0]

	if word == "min" || word == "Min" || word == "MIN" {
		if len(numbers) == 0 {
			return 0
		}

		for _, i := range numbers {
			if i < min {
				min = i
			}
		}

	} else if word == "max" || word == "Max" || word == "MAX" {
		if len(numbers) == 0 {
			return 0
		}

		for _, i := range numbers {
			if i > max {
				max = i
			}
		}
		return max
	}
	if word == "min" || word == "Min" || word == "MIN" {
		return min
	}
	return max
}
