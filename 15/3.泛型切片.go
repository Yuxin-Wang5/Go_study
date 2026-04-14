package main

type MySlice[T any] []T

func main() {
	var mySlice MySlice[string]
	mySlice = append(mySlice, "	yaoyao")
	var intSlice MySlice[int]
	intSlice = append(intSlice, 2)
}
