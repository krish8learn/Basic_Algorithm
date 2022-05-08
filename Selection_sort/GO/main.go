package main

import (
	"fmt"
)

func main() {
	fmt.Println((SelectionSort([]int{4, 3, 2, 1, 0, -1, -99})))
}

func SelectionSort(A []int) []int {
	//result := make([]int, len(A))
	for i := 0; i < len(A)-1; i++ {
		swapIndex := SubarrayMin(A, i)
		Swap(A, i, swapIndex)
	}
	return A
}

func SubarrayMin(A []int, startIndex int) int {
	minIndex := startIndex
	minvalue := A[startIndex]

	for i := startIndex + 1; i < len(A); i++ {
		if A[i] < minvalue {
			minIndex = i
			minvalue = A[i]
		}
	}
	return minIndex
}

func Swap(A []int, firstIndex, secondIndex int) {
	temp := A[firstIndex]
	A[firstIndex] = A[secondIndex]
	A[secondIndex] = temp
	return
}
