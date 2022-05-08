package main

import "fmt"

func main() {
	fmt.Println((InsertionSort([]int{4, 3, 2, 1, 0, -1, -99})))
}

func insert(A []int, rightIndex, value int) {
	endpos := 0
	for i := rightIndex; i < len(A); i++ {
		if A[i] == value {
			endpos = i
			break
		}
	}

	for i := 0; i <= rightIndex; i++ {
		if value < A[i] {
			for j := endpos; j > i; j-- {
				k := j
				k -= 1
				A[j] = A[k]
				//System.out.println(array[j]+","+j+","+k+","+array[k]);
			}
			A[i] = value
			break
		}
	}

}

func InsertionSort(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		insert(A, i+1, A[i+1])
	}
	return A
}
