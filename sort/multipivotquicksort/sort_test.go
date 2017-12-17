package multipivotquicksort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestOnePivot(t *testing.T) {
	a := []int{5, 3, 4, 7, 1, 9, 3, 4, 9, 1, 4, 7}
	res, err := MultiPivot(a, 1, true)
	if err != nil {
		t.Error(err)
	}
	if isOrdered(res) == false {
		t.Errorf("%v is not ordered", a)
	}
}

func Test3Pivot(t *testing.T) {
	a := []int{5, 3, 4, 7, 1, 9, 3, 4, 9, 1, 4, 7}
	res, err := MultiPivot(a, 3, true)
	if err != nil {
		t.Error(err)
	}
	if isOrdered(res) == false {
		t.Errorf("%v is not ordered", a)
	}
}
func Test3PivotAsync(t *testing.T) {
	a := []int{5, 3, 4, 7, 1, 9, 3, 4, 9, 1, 4, 7}
	res, err := MultiPivot(a, 3, false)
	if err != nil {
		t.Error(err)
	}
	if isOrdered(res) == false {
		t.Errorf("%v is not ordered", a)
	}
}

func TestMultiPivotTable(t *testing.T) {
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
		res, err := MultiPivot(orig, test.pivots, true)
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

//generator predictable array generator
func generator(seed int64, max, count int) (res []int) {
	res = make([]int, count)
	rand.Seed(seed)
	for i := range res {
		res[i] = rand.Intn(max)
	}
	return
}

func BenchmarkFivePivot5P300000L(b *testing.B) {
	orig := generator(2, 10000, 300000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(orig))
		copy(arr, orig)
		MultiPivot(arr, 5, true)
	}
}

func BenchmarkFivePivotAsync2P300000L(b *testing.B) {
	orig := generator(2, 10000, 300000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(orig))
		copy(arr, orig)
		MultiPivot(arr, 2, false)
	}
}

func BenchmarkFivePivotAsync5P300000L(b *testing.B) {
	orig := generator(2, 10000, 300000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(orig))
		copy(arr, orig)
		MultiPivot(arr, 5, false)
	}
}

func BenchmarkFivePivotAsync7P300000L(b *testing.B) {
	orig := generator(2, 10000, 300000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(orig))
		copy(arr, orig)
		MultiPivot(arr, 7, false)
	}
}

func ExampleMultiPivot() {
	a := []int{5, 3, 4, 7, 1, 9, 3, 4, 9, 1, 4, 7}
	var pivotCount uint8 = 3
	singleThread := false
	res, err := MultiPivot(a, pivotCount, singleThread)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	fmt.Printf("Result is\n%v", res)
	//Output: Result is
	//[1 1 3 3 4 4 4 5 7 7 9 9]
}
