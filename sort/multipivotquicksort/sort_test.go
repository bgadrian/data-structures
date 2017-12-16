package multipivotquicksort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSwap(t *testing.T) {
	a := []int{3, 2, 1}
	swap(a, 0, 2)
	if a[0] != 1 && a[2] != 3 {
		t.Errorf("swap failed %v",
			a)
	}
}

func TestFivePivot(t *testing.T) {
	a := []int{5, 3, 4, 7, 1, 9, 3, 4, 9, 1, 4, 7}
	res, err := FivePivot(a, 1)
	if err != nil {
		t.Error(err)
	}
	if isOrdered(res) == false {
		t.Errorf("%v is not ordered", a)
	}
}

func TestFivePivotRandom(t *testing.T) {
	type one struct {
		seed   int64
		max    int
		count  int
		pivots uint8
	}
	table := []*one{
		{1, 100, 10, 1},
		{1, 100, 10, 12},
		{2, 1000, 10, 2},
		{2, 1000, 10, 3},
		{2, 1000, 10, 4},
		{2, 1000, 10, 2},
		{3, 5000, 10, 3},
		{4, 100000, 10, 4},
		{5, 100, 30, 1},
		{1, 100, 10, 5},
		{2, 1000, 10, 5},
		{3, 5000, 10, 5},
		{4, 100000, 10, 5},
		{5, 100, 30, 5},
	}
	for _, test := range table {
		defer func() {
			if r := recover(); r != nil {
				fmt.Errorf("%v %v", test, r)
			}
		}()

		orig := generator(test.seed, test.max, test.count)
		res, err := FivePivot(orig, test.pivots)
		if err != nil {
			t.Error(err)
		}
		if isOrdered(res) == false {
			t.Errorf("%v \norig:%v \nsorted:%vis not ordered", test, orig, res)
		}
		if len(res) != test.count {
			t.Errorf("wrong count, is %v want %v",
				len(res), test.count)
		}
	}
}

func isOrdered(arr []int) bool {
	for i, v := range arr {
		if i == 0 {
			continue
		}
		if arr[i-1] > v {
			return false
		}
	}
	return true
}

func generator(seed int64, max, count int) (res []int) {
	res = make([]int, count)
	rand.Seed(seed)
	for i := range res {
		res[i] = rand.Intn(max)
	}
	return
}
