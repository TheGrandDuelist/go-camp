package main

import (
	"fmt"
	"reflect"
	"strings"
)

func map_alloc() {
	m := make(map[int]int)
	//m := make(map[int]int, 200)
	for i := 0; i < 200; i++ {
		m[i] = i
	}
}

func slice_alloc() {
	//s := make([]int, 0)
	s := make([]int, 100)
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
}

func test_string() {
	s := "   "
	s = strings.TrimSpace(s)
	s = "   111   "
	s = strings.TrimRight(s, " ")
	fmt.Println(s)
	if s == "" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func test_map() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	if _, ok := m["c"]; !ok {
		fmt.Println("m[\"ok\"] not exist")
	}
}

func testStringModify() {
	s := "abc"
	m, n := &s, &s
	*m = "efg"
	s = "mng"
	fmt.Println(m, *m, n, *n)
	fmt.Println(s)
}

func testEqual() {
	a := make([]int, 20, 100)
	//a[0] = 1
	b := make([]int, 10, 50)

	fmt.Println(reflect.DeepEqual(a, b))

	c := "str"
	d := "Str"
	fmt.Println(strings.ToUpper(d) == c)
}

func testPanic() {
	defer func() {
		fmt.Println("recovered ", recover())
	}()
	panic("not good")
}

func testCap() {
	a := make([]int, 2, 100)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 2, 2)
	b = append(b, 1)
	b = append(b, 1)
	b = append(b, 1)
}

func testFor() {
	a := []string{"one", "two", "three"}
	for _, v := range a {
		go func() {
			b := v
			fmt.Println(b)
		}()
	}

	var i = 1
	defer fmt.Println(func() int { return i * 2 }())
	i = 2
}

/*func main() {
	testFor()
	time.Sleep(time.Second * 3)
}*/
