package main

import "fmt"

func main() {

}

func slices() {
	// arr => [n]t{values ...}; n = size, t - type
	// например, UUID:258d5bc0-d6c1-489a-b933-e3486ff772b9 -> 16 bytes

	/* Из чего состоит slice:
	*1)
		!Ссылка на некий базовый массив -> pointer: *array
	*2)
		!Ссылка указывает на первый элемент массива, а последний элемент
		!определяется с помщью параметра lenght: int
	*3)
		!capacity - количество элементов базового массива
	*/

	myArray := [13]string{"", "a", "b", "c", "d", "f", "f", "@", "#", "z", "x", "yx", "cc"}
	slice1 := myArray[4:7] // -> len(slice1) = 3, тк [4:7),
	//cap(slice1) = 9,
	//указатель смотрит на элемент с номером 4
	fmt.Println(len(slice1), cap(slice1), slice1[0])

	slice2 := myArray[6:9] // -> len(slice2) = 3, тк [6:9),
	//cap(slice2) = 7,
	//указатель смотрит на элемент с номером 6
	fmt.Println(len(slice2), cap(slice2), slice2[0])

	/* Добавление элементов в slice
	   *slice2 = ["f", "f", "@", "#", "z", "x"]
			    !|____________|
					!len = 3
		 		?|___________________________|
		 			    ?cap = 6

		Как выгладит работа append:
		slice2 = append(slice2, "Q", "W", "E")
	   *slice2 = ["f", "f", "@", "Q", "W", "E"]
			    !|____________________________|
						!len = 6
		 		?|___________________________|
		 			    ?cap = 6

		Что будет после вызова append, когда len = cap?
		=> Будет создан новый базовый массив и в большинстве случаев
		 его длинна в 2 раза больше изначального.
	*/
}

func zeroSliceValue() {
	// НУЛЕВОЕ ЗНАЧЕНИЕ СЛАЙСА

	var list []int
	fmt.Println("без len:", list == nil) // => true, создали переменную list типа slice, но ничего не положили в нее
	fmt.Println("через len:", len(list) == 0)

	list = []int{}
	fmt.Println("без len:", list == nil) // => false, создали переменную list типа slice и проинициализировали ее, пусть и пустую, то есть какое-то значение у нее уже есть
	fmt.Println("через len:", len(list) == 0)
	// !лучше делать через LEN()!
}

func allocationOfMemoryForSlice() {
	// АЛЛОКАЦИЯ ПАМЯТИ ДЛЯ СЛАЙСА
	sl := make([]int, 5, 10)
	fmt.Println(len(sl), cap(sl))

	sl_test := []int{1, 2, 3, 4, 8}
	fmt.Println(double(sl_test))

}

func passingASliceByValue() {
	// ПЕРЕДАЧА СЛАЙСА ПО ЗНАЧЕНИЮ
	sample_slice := []int{1, 2, 3, 4}
	fmt.Println("before:", sample_slice[:1])
	handle(sample_slice)
	fmt.Println("after:", sample_slice)

}

func handle(list []int) {
	//list[1] = 123
	_ = append(list, 5)
	fmt.Println("append:", list)
}

func double(nums []int) []int {
	//var res []int, если создавать таким образом, то длина = 1, 2, 4, 4, 8, что не очень оптимально
	//res := make([]int, 0, len(nums))
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = num * 2
	}
	return res
}
