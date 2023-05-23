package main

import (
	"fmt"
)

type MyInt int
type NewInt = int

type person struct {
	name, city, gender string
	age                int
}

type student struct {
	name string
	age  int
}

type Dog2 struct {
	string
	int
}

func entry_16() {
	var a MyInt
	var b NewInt

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	var c person
	c.name = "xiaoC"
	c.city = "shanghai"
	c.gender = "male"
	c.age = 10
	fmt.Println("p =", c, c.name, c.age)
	fmt.Printf("p = %v\n", c)
	fmt.Printf("p = %#v\n", c)

	var d struct {
		Name string
		Age  int
	}
	d.Name = "xiaoD"
	d.Age = 20
	fmt.Println(d)

	var f *person = new(person)
	f.city = "nanjing"
	f.age = 20
	f.name = "xiaoF"
	fmt.Println(f)
	fmt.Printf("p=%v\n", f)
	fmt.Printf("p=%#v\n", f)

	var g = &person{}
	g.age = 30
	g.gender = "male"
	g.city = "test"
	g.name = "xiaoG"
	fmt.Printf("g=%#v\n", g)

	p2 := person{
		age:    10,
		name:   "p2",
		gender: "male",
		city:   "beijing",
	}

	fmt.Printf("p2 = %#v\n", p2)

	p3 := &person{
		age:    10,
		name:   "p2",
		gender: "male",
	}
	fmt.Printf("p3 = %#v\n", p3)

	p4 := &person{
		"p4",
		"abc",
		"female",
		10,
	}

	fmt.Printf("p4 = %#v\n", p4)

	m := make(map[string]*student)
	stus := []student{
		{name: "stu1", age: 10},
		{name: "stu2", age: 20},
		{name: "stu3", age: 22},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v)
	}

	p4.Dream()

	p5 := NewPerson("p5", 10)
	p5.Dream()
	p5.SetAge(22)
	p5.Dream()

	p6 := NewPerson("p6", 26)
	p7 := NewPerson("p7", 27)

	p6.SetAge(266)
	p6.Dream()
	p7.SetAge2(277)
	p7.Dream()

	d1 := Dog2{
		"dog1", 2,
	}
	fmt.Println(d1, d1.string, d1.int)
}

func NewPerson(name string, age int) *person {
	return &person{
		name:   name,
		age:    age,
		city:   "shanghai",
		gender: "male",
	}
}

func (p person) Dream() {
	fmt.Printf("%s %d dreams learn golang well!\n", p.name, p.age)
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p person) SetAge2(age int) {
	p.age = age
}

func (a MyInt) test() {
	fmt.Println(a, " is int")
}
