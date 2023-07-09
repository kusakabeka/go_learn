package main

import "fmt"

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

/* -- просто функция --
func printUserInfo(user User) {
	fmt.Println(user.name, user.age, user.sex, user.weight, user.height)
}
*/

// -- метод структыры -- (данные просто копируются)
func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u User) GetName() string {
	return u.name
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
		age:    Age(age),
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
	user1.SetName("Sergey")
	fmt.Println(user1.GetName())
	fmt.Println(user2.GetName())

	fmt.Println(user1.age.isAdult())

}
