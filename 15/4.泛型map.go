package main

import "fmt"

type MyMap[K string | int, V any] map[K]V

func main() {
	var myMap = make(MyMap[string, string])
	myMap["name"] = "奇异"
	fmt.Println(myMap)
}
