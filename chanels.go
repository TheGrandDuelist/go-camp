package main

import "fmt"

func entry_2() {
	var ch chan int = make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("send success")

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("send ", i)
		}
		close(c)
	}()

	for i := 0; i < 5; i++ {
		if data, ok := <-c; ok {
			fmt.Println("recv data ", data)
		} else {
			break
		}
	}

	fmt.Println("main 结束")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			data, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- data * data
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("receive success ", ret)
}
