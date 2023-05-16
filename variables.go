package main

import (
	"fmt"
	"time"
)

const s string = "constant"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)

	const n = 200
	const m = "abc"

	fmt.Println(n, m)
	fmt.Println("------------------------")
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("------------------------")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("------------------------")
	for n := 0; n < 20; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("------------------------")

	if 7%2 == 0 {
		fmt.Println("7/2 == 0")
	} else {
		fmt.Println("7/2 == ", 7/2)
	}

	if num := 9; num < 0 {
		fmt.Println(num)
	} else if num < 5 {
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
	fmt.Println("------------------------")

	ii := 2
	fmt.Print("Write ", ii, " as ")
	switch ii {
	case 1:
		fmt.Println("one")
		break
	case 2:
		fmt.Println("two")
		break
	}
	fmt.Println("------------------------")
	line()
	switch time.Now().Weekday() {
	case time.Monday:
	case time.Tuesday:
		break
	case time.Wednesday:
		break
	case time.Thursday:
		break
	case time.Friday:
		break
	case time.Saturday:
		break
	case time.Sunday:
		break
	}
	fmt.Println(time.Monday)
	line()
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("beforenoon")
	default:
		fmt.Println("afternoon")
	}
	fmt.Println(t.Hour())
	line()
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("unknown ", t)
		}
	}
	whatAmI(false)
	whatAmI(56)
	whatAmI("abc")
	line()
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}
	line()
	x, y := foo(1)
	fmt.Println(x, y)

	const (
		a1 = "1123"
		b5
		b1 = iota
		c1
		b2
		c3 = false
	)
}

func line() {
	fmt.Println("------------------------")
}

func foo(a int) (int, int) {
	fmt.Println("test ", a)
	return 1, 2
}
