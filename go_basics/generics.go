package main

import "fmt"

func main() {

	a := []int64{1, 2, 3, 4, 5, 6}
	b := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	fmt.Println(sumOfInt64(a))
	fmt.Println(sumOfFloat64(b))

}

func sumOfInt64(input []int64) int64 {
	sum := int64(0)
	for _, v := range input {
		sum += v
	}
	return sum
}

func sumOfFloat64(input []float64) float64 {
	sum := float64(0)
	for _, v := range input {
		sum += v
	}
	return sum
}
