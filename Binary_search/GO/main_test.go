package main

import "testing"

type caseStruct struct{
	field1 []int
	field2 int
	expect int
}
func TestBinarySearch(t *testing.T) {
	cases := []caseStruct{
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},3,1},
		{[]int{22, 44, 66, 88},88,3},
		{[]int{99,101},90,-1},
	}

	for _, val := range cases{
		if output:= BinarySearch(val.field1,val.field2); output != val.expect{
			t.Error("failed")
		}
	}

}
//test command go test or go test -v