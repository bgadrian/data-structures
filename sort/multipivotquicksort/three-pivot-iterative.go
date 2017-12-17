package multipivotquicksort

func ThreePivot(A []int, left, right int) {
	a := left + 2
	b := left + 2
	c := right - 1
	d := right - 1

	pivot1 := A[left]
	pivot2 := A[left+1]
	pivot3 := A[right]

	for b <= c {
		for A[b] < pivot2 && b <= c {
			if A[b] < pivot1 {
				swap(A, a, b)
				a = a + 1
			}
			b = b + 1
		}

		for A[c] > pivot2 && b <= c {
			if A[c] > pivot3 {
				swap(A, c, d)
				d = d - 1
			}
			c = c - 1
		}

		if b <= c {
			if A[b] > pivot3 {
				if A[c] < pivot1 {
					swap(A, b, a)
					swap(A, a, c)
					a = a + 1
				} else {
					swap(A, b, c)
				}
				swap(A, c, d)
				b = b + 1
				c = c - 1
				d = d - 1
			} else {
				if A[c] < pivot1 {
					swap(A, b, a)
					swap(A, a, c)
					a = a + 1
				} else {
					swap(A, b, c)
				}
				b = b + 1
				c = c - 1
			}
		}
	}
	a = a - 1
	b = b - 1
	c = c + 1
	d = d + 1
	swap(A, left+1, a)
	swap(A, a, b)
	a = a - 1
	swap(A, left, a)
	swap(A, right, d)
}

func swap(A []int, a, b int) {
	A[a], A[b] = A[b], A[a]
}
