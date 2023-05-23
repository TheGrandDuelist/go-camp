package main

import (
	"fmt"
	"log"
	"time"
)

func testTimeFunc() {
	later := time.Now().Add(time.Second * 60)
	fmt.Println(later.Year(), later.Month(), later.Day(), later.Minute())

	/*ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}*/

	now := time.Now()
	fmt.Println(now.Format("2006-01-02 03:04"))
	log.Println("log test1")
	a := 0
	if a == 0 {
		panic("a is o")
		return
	}
	log.Println("log test2")
}

func testTime() {
	today := time.Date(2023, 5, 31, 0, 0, 0, 0, time.Local)
	d := today.Day()
	fmt.Println(d)
	nextToday := today.AddDate(0, 1, 0)
	fmt.Println(nextToday.Format("200601020304"))

	nToday := today.AddDate(0, 2, -d)
	fmt.Println(nToday)
}

func main() {
	testTime()
}
