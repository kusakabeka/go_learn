package main

import "fmt"

func main() {
	users := map[string]int{
		"name_1": 15,
		"name_2": 24,
		"name_3": 55,
		"name_4": 5,
	}
	// fmt.Println(user["name_5"]) -> return 0
	age, ok := users["name_1"]
	fmt.Println(age, ok)

	users["name_5"] = 21

	delete(users, "name_2")
	for key, value := range users {
		fmt.Println(key, value)
	}

}
