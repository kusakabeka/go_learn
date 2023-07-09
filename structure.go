package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// -- метод структыры -- (данные просто копируются)
func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

func (u *User) printUserInfoWithPointer(name string) {
	u.name = name
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

// -- конструстор --
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

	user1.printUserInfo()
	user2.printUserInfo()
	user3.printUserInfo()
	user4.printUserInfo()

	//fmt.Println(user1.name)

	user1.printUserInfoWithPointer("Oleg")
	fmt.Println(user1)

}

/* -- просто функция --
func printUserInfo(user User) {
	fmt.Println(user.name, user.age, user.sex, user.weight, user.height)
}
*/
