package tree

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
