package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scoreMap := make(map[string]int, 2)
	scoreMap["a"] = 10
	scoreMap["c"] = 30
	scoreMap["b"] = 20

	fmt.Println(scoreMap)
	fmt.Println(scoreMap["c"])

	v, ok := scoreMap["d"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not exists")
	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	for k, _ := range scoreMap {
		fmt.Println(k)
	}

	fmt.Println(scoreMap)
	delete(scoreMap, "a")
	fmt.Println(scoreMap)
	delete(scoreMap, "f")
	fmt.Println(scoreMap)

	stuMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		stuMap[key] = value
	}
	fmt.Println(stuMap)
	fmt.Println("-------------------")

	var keys = make([]string, 0, 200)
	for k := range stuMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, stuMap[key])
	}
	fmt.Println("-------------------")
	for k, v := range stuMap {
		fmt.Println(k, v)
	}
	fmt.Println("-------------------")

	var sliceMap = make(map[string][]string, 3)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 2)
	}

	value = append(value, "上海", "北京")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
