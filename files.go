package main

import (
	"fmt"
	"os"
)

func readFiles() {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	readFiles()
}
