package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	testArr := []int{1, 2, 3, 7, 4, 5, 6}
	testCases := []struct{
		in int
		want bool
	}{
		{1, true},
		{0 , false},
	}
	for _, tc := range testCases {
		result := BinarySearch(testArr, tc.in)
		if result != tc.want {
			t.Fail()
		}
	}
}
