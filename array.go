package main

import (
	"fmt"
	"go/types"
)

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	a[1] = 21
	fmt.Println(a)
	var s1 []int
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 not nil")
	}

	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s2, s3)

	var s4 = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case types.Array:
			fmt.Println("is array")
		case types.Slice:
			fmt.Println("is Slice")
		default:
			fmt.Println("unknown")
		}
	}
	whatAmI(s5)

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[2:8]
	fmt.Println(slice0)
	var slice1 = arr[0:6]
	fmt.Println(slice1)
	var slice2 = arr[:]
	fmt.Println(slice2)
	var slice3 = arr[:5]
	fmt.Println(slice3)
	var slice4 = make([]int, 2, 5)
	fmt.Println(slice4)

	slice5 := arr[2:4]
	slice5[0] += 100
	slice5[1] += 200
	fmt.Println(slice5)
	fmt.Println(arr)

	s := []int{1, 2, 3}
	p := &s[2]
	*p += 100
	fmt.Println(s)

	slice6 := arr[:2:4]
	fmt.Println(slice6, len(slice6), cap(slice6))
	fmt.Println(arr)
	slice6 = append(slice6, 1, 1, 1)
	fmt.Println(slice6, arr)

	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	ss1 := arr2[6:]
	ss2 := arr2[:5]
	fmt.Println(ss1, ss2)
	copy(ss2, ss1)
	fmt.Println(ss1, ss2)

	for index, value := range ss2 {
		fmt.Println(index, value)
	}

	str := "hello world"
	fmt.Println(&str)
	str = "1231"
	fmt.Println(&str)
}
