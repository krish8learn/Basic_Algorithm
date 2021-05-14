package main

import "fmt"

func main() {
	list1 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	fmt.Println(BinarySearch(list1, 3))

	list2 := []int{22, 44, 66, 88}
	fmt.Println(BinarySearch(list2, 88))

	list3 := []int{99, 100}
	fmt.Println(BinarySearch(list3, 89))
}

func BinarySearch(elements []int, targetValue int) int {
	min := 0
	max := len(elements) - 1

	for max >= min {
		guess := min + (max-min)/2

		if elements[guess] == targetValue {
			return guess
		} else if elements[guess] > targetValue {
			max = guess - 1
		} else if elements[guess] < targetValue {
			min = guess + 1
		}
	}
	return -1
}
