package heap

import "testing"

func BenchmarkIHMinBuilding1000N100P(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewImplicitHeapMin(false)
		addIHNodesWithMaxP(h, 1000, 100)
	}
}

func BenchmarkIHMaxBuilding1000N100P(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewImplicitHeapMax(false)
		addIHNodesWithMaxP(h, 1000, 100)
	}
}

func BenchmarkIHMinBuilding1000000N255P(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewImplicitHeapMin(false)
		addIHNodesWithMaxP(h, 1000000, 255)
	}
}

func BenchmarkIHMin1000N2Ops255P(b *testing.B) {
	h := NewImplicitHeapMin(false)
	addIHNodesWithMaxP(h, 1000, 255)
	benchmarkIHOneOps(h, b, 255)
}

func BenchmarkIHMin100000N2Ops255P(b *testing.B) {
	h := NewImplicitHeapMin(false)
	addIHNodesWithMaxP(h, 100000, 255)
	benchmarkIHOneOps(h, b, 255)
}

func BenchmarkIHMin1000000N2Ops255P(b *testing.B) {
	h := NewImplicitHeapMin(false)
	addIHNodesWithMaxP(h, 1000000, 255)
	benchmarkIHOneOps(h, b, 255)
}

func benchmarkIHOneOps(h ImplicitHeap, b *testing.B, maxP int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(i%maxP, i)
		_, ok := h.Pop()
		if ok == false {
			b.Error(ok)
			break
		}
	}
}

func BenchmarkIHMin1000N1P(b *testing.B) {
	h := NewImplicitHeapMin(false)
	benchmarkIHConstant(h, 1000, b)
}

func BenchmarkIHMin1000000N1P(b *testing.B) {
	h := NewImplicitHeapMin(false)
	benchmarkIHConstant(h, 1000000, b)
}

func BenchmarkIHMin10000000N1P(b *testing.B) {
	h := NewImplicitHeapMin(false)
	benchmarkIHConstant(h, 10000000, b)
}

//benchmarkIHConstant Test that when 1 priority is used, the complexity should be O(1)
//need this behaviour when used by Hierarchical Heap and buckets == count(priorities)
func benchmarkIHConstant(h ImplicitHeap, count int, b *testing.B) {
	for i := 0; i < count; i++ {
		h.Push(2, 2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(2, 2)
		_, ok := h.Pop()
		if ok == false {
			b.Error(ok)
			break
		}
	}
}
