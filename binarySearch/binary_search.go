package binarySearch

import "math"

func BinarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		m := int(math.Floor(float64(lo + (hi - lo)/2)))
		v := haystack[m]

		if v == needle {
			return true
		} else if v > needle {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return false
}
