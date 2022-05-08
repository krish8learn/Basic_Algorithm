package main

import "testing"

type caseStruct struct {
	field1 []int
	field2 []int
}

func TestInsertionSort(t *testing.T) {
	cases := []caseStruct{
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 2, 1, 0, -1, -99}, []int{-99, -1, 0, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, val := range cases {
		output := InsertionSort(val.field1)
		for i := 0; i < len(val.field1); i++ {
			if output[i] != val.field2[i] {
				t.Error("Failed")
				break
			}
		}
	}
}
