package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, a := range slice {
		slice[i] = f(a)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	for i, a := range array {
		array[i] = f(a)
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}

	mapSlice(addOne, intsSlice)
	fmt.Println("Splice after addOne:", intsSlice)

	intsArray := [5]int{1, 2, 3, 4, 5}
	//intsArray = mapArray(addOne, intsArray)
	fmt.Println("Array after addOne:", intsArray)

	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println("newSlice:", newSlice, "intsSlice:", intsSlice)
}
