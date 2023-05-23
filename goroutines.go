package main

import (
	"fmt"
	"runtime"
)

func entry_10() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}

	var ch chan int
	fmt.Println(ch)
	ch <- 10
	x := <-ch
	fmt.Println(x)
}
