package multipivotquicksort

import "testing"

func TestSwap(t *testing.T) {
	a := []int{3, 2, 1}
	swap(a, 0, 2)
	if a[0] != 1 && a[2] != 3 {
		t.Errorf("swap failed %v",
			a)
	}
}

func TestThreePivot(t *testing.T) {
	arr := []int{5, 3, 4, 7, 1, 9, 3}
	left := 1
	right := 4
	ThreePivot(arr, left, right)

	if isOrdered(arr) == false {
		t.Errorf("Result is not sorted %v", arr)
	}
}
