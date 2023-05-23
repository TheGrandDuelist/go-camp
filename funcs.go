package main

import (
	"fmt"
)

func entry_7() {
	var a, b = 1, 2
	fmt.Println(a, b)
	//swap(&a, &b)
	swap2(a, b)
	fmt.Println(a, b)

	fmt.Println("sum 1->5= ", sum(1, 2, 3, 4, 5))
	fmt.Println("sum 1->10= ", sum([]int{1, 2, 3}...))

	//print_tt(1, 2, "test", "a")

	d := test(1, 2)
	fmt.Println(d)

	add(1, 2)

	fn := func() {
		fmt.Println("hello, world!")
	}
	fn()

	fns := []func(x int) int{
		func(x int) int {
			return x + 1
		},

		func(x int) int {
			return x + 2
		},
	}

	fmt.Println(fns[1](100))

	s := struct {
		fn func() string
		fm func()
	}{
		fn: func() string {
			return "hello, world!"
		},
		fm: func() {
			fmt.Println("fm")
		},
	}
	fmt.Println(s.fn())
	s.fm()
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func swap2(x, y int) {
	temp := x
	x = y
	y = temp
}

func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

/*func print_tt(args ...interface{}) {
	fmt.Println(args)
}*/

func test(a, b int) (c int) {
	c = a + b
	return
}

func add(x, y int) (z int) {
	defer func() {
		fmt.Println("defer ", z)
	}()
	z = x + y
	fmt.Println("main ", z)
	return z + 200
}
