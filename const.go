package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a = iota
	b = (1 + iota)
	_
	d
)

func main() {
	c := [5]int{2: 100, 4: 200}
	fmt.Println(c)
	d := [...]struct {
		name string
		age  uint
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(d)
}
