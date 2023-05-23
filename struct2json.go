package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Gender string
}

type Class struct {
	Title    string     `json:"-"`
	Students []*Student `json:"students_test"`
}

func entry_15() {
	c := &Class{
		Title:    "102",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 100; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Id:     i,
			Gender: "female",
		}

		c.Students = append(c.Students, stu)
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Printf("%s\n", data)
	str := `{"Title":"15","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("unjson err")
		return
	}
	fmt.Printf("%#v\n", c1)

}
