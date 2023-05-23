package main

import "fmt"

type Student2 struct {
	Name string
	Age  int
}

func entry_8() {
	var stu2 = Student2{
		Name: "lili",
		Age:  20,
	}

	study(stu2)
	fmt.Printf("%#v\n", stu2)

	study2(&stu2)
	fmt.Printf("%#v\n", stu2)

	stu2.study3()
	fmt.Printf("%#v\n", stu2)
	stu2.study4()
	fmt.Printf("%#v\n", stu2)
}

func study(stu Student2) {
	stu.Name = "lili study"
}

func study2(stu *Student2) {
	stu.Name = "lili2 study"
}

func (stu Student2) study3() {
	stu.Name = "zhangsan"
	fmt.Println(stu.Name, "study3")
}

func (stu *Student2) study4() {
	stu.Name = "lisi"
	fmt.Println(stu.Name, "study4")
}
