package main

import "fmt"

func entry_3() {
	c := funca()
	c()
	c()
	c()
	funca()

	f1, f2 := test01(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))

	fmt.Printf("%d != %d\n", 6, factorial(6))
	fmt.Printf("%d != %d\n", -1, factorial(-1))

	for i := 0; i < 10; i++ {
		fmt.Println(fabonacci(i))
	}
}

func funca() func() int {
	var i = 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func test01(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fabonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fabonacci(i-1) + fabonacci(i-2)
}
