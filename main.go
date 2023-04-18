package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	}
	return data
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = AddElement(arr, 6, "up")
	fmt.Println(arr)

	arr2 := []int{1, 2, 3, 4, 5}
	arr2 = AddElement(arr2, 6, "down")
	fmt.Println(arr2)
}
