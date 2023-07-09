package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

// конструстор
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	//user := struct {
	//	name   string
	//	age    int
	//	sex    string
	//	weight int
	//	height int
	//}
	user1 := User{"ivan", 23, "Male", 75, 185}
	user2 := User{"anton", 33, "Male", 75, 140}
	user3 := User{"anna", 27, "Female", 75, 176}
	user4 := NewUser("fedot", "male", 123, 31, 144)

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	fmt.Println(user4)
	fmt.Println(user1.name)
}

func printUserInfo(user User) {

}
