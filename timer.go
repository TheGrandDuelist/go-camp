package main

import (
	"fmt"
	"time"
)

func entry_19() {
	timer1 := time.NewTimer(time.Second * 5)
	fmt.Println(time.Now().UnixMilli())
	for {
		<-timer1.C
		fmt.Println("time on")
	}

	for {

	}
}
