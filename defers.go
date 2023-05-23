package main

import "fmt"

func entry_6() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}

	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}
