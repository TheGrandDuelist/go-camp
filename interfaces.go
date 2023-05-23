package main

import (
	"encoding/json"
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student3 struct{}

func (stu *Student3) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func entry_12() {
	var peo People = &Student3{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	var stuInfo = make(map[string]interface{})
	stuInfo["name"] = "zhangsna"
	stuInfo["age"] = 10
	stuInfo["gender"] = "female"
	stuInfo["grade"] = false
	fmt.Printf("%#v\n", stuInfo)

	str, err := json.Marshal(stuInfo)
	if err != nil {
		fmt.Println("json erro")
		return
	}
	fmt.Println(string(str))

}
