package main

import "fmt"

func main() {
	user := map[string]int{
		"name_1": 15,
		"name_2": 24,
		"name_3": 55,
		"name_4": 5,
	}
	fmt.Println(user)
	// users :=[]string{"1", "2"}
	//users := make(map[string]int)
	//user["Ivan"] = 12
	//for key, value := range users {
	//	fmt.Println(key, value)
	//}

	// fmt.Println(user[""]) -> return 0
	age, ok := user["name_1"]
	fmt.Println(age, ok)

	user["name_5"] = 21

	delete(user, "name_2")

	fmt.Println(len(user))
	// cap() - не работает с мапами

}
