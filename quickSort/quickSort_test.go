package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	QuickSort(&arr)

	if !reflect.DeepEqual(arr, []int{3, 4, 7, 9, 42, 69, 420}) {
		t.Fail()
	}
}
