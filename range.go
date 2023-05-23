package main

import "fmt"

func entry_14() {
	str := "abc"
	for _, c := range str {
		fmt.Printf("%c\n", c)
	}

	for i := range str {
		fmt.Println(str[i])
	}

	arr := [3]int{0, 1, 2}

	for i, v := range arr {
		if i == 0 {
			arr[i] = 999
			arr[1], arr[2] = 999, 999
		}
		println(i, v)
		//arr[i] = v + 100
	}
	fmt.Println(arr)

	s := []int{0, 1, 2, 3, 4}

	for i, v := range s {
		if i == 0 {
			s := s[:3]
			s[2] = 99
		}

		fmt.Println(i, v)
	}

	fmt.Println(s)

}
