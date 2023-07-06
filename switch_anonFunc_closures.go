package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := prediction("пт")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
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
