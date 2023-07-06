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

	fmt.Println(findMinOrMax("min", 1, 22, 34, 51, 25))
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
	if word == "min" || word == "Min" || word == "MIN" {
		if len(numbers) == 0 {
			return 0
		}
		min := numbers[0]

		for _, i := range numbers {
			if i < min {
				min = i
			}
		}
		return min
	} else if word == "max" || word == "Max" || word == "MAX" {
		if len(numbers) == 0 {
			return 0
		}
		max := numbers[0]

		for _, i := range numbers {
			if i > max {
				max = i
			}
		}
		return max
	}
	return 0
}
