package main

import "fmt"

type Address struct {
	Province string
	City     string
}

type User struct {
	Username string
	Password string
	Address  Address
}

type User2 struct {
	Username  string
	Password  string
	Favorites []string
}

type User3 struct {
	Name   string
	Gender string
	Address
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s can move\n", a.name)
}

type Dog struct {
	Feet int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s can wang\n", d.name)
}

func entry_17() {
	add := Address{
		City: "tess", Province: "tset",
	}

	fmt.Println(add)
	user1 := User{
		Username: "admin",
		Password: "123456",
		Address:  Address{Province: "shanghai", City: "shanghai"},
	}

	fmt.Println(user1)
	fmt.Printf("%#v\n", user1)

	user2 := User2{
		"test1", "test2",
		[]string{"test1", "test2"},
	}

	fmt.Println(user2)

	var user3 User3
	user3.Gender = "male"
	user3.Name = "user3"
	user3.Address.Province = "chongqing"
	user3.City = "chongqing"
	fmt.Printf("%#v\n", user3)

	dog := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "lele",
		},
	}

	dog.move()
	dog.wang()
}
