package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(&array[0])
	array[0] = 20
	fmt.Println(&array[0])
}
