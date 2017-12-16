/*Package multipivotquicksort
http://epubs.siam.org/doi/pdf/10.1137/1.9781611973198.6
https://web.archive.org/web/20151002230717/http://iaroslavski.narod.ru/quicksort/DualPivotQuicksort.pdf
http://cs.stanford.edu/~rishig/courses/ref/l11a.pdf
http://rerun.me/2013/06/13/quicksorting-3-way-and-dual-pivot/
https://lusy.fri.uni-lj.si/sites/lusy.fri.uni-lj.si/files/publications/alopez2014-seminar-qsort.pdf
http://eiche.theoinf.tu-ilmenau.de/quicksort-experiments/uploads/multi-pivot-paper-draft-2015-10-12.pdf
http://iopscience.iop.org/article/10.1088/1757-899X/180/1/012051/pdf

"First, we have confirmed previous experimental results
on Yaroslavskiy’s dual-pivot algorithm under a basic en-
vironment thus showing that the improvements are not
due to JVM side effects.  We designed and analysed a
3-pivot approach to quicksort which yielded better re-
sults both in theory and in practice.  Furthermore, we
provided strong evidence that much of the runtime im-
provements are from cache effects in modern architec-
ture by analysing cache behavior.
We have learned that due to the rapid development
of  hardware,  many  of  the  results  from  more  than  a
decade ago no longer hold."
*/
package multipivotquicksort

import (
	"fmt"
	"sort"
)

func swap(A []int, a, b int) {
	A[a], A[b] = A[b], A[a]
}

func FivePivot(list []int, pivotCount uint8) (result []int, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("%v %v",
				r, list)
		}
	}()
	n := len(list)
	if n <= 1 {
		result = list
		return
	}

	if pivotCount <= 1 {
		pivotCount = 1
	}

	if n <= 20 || n/int(pivotCount) < 10 {
		sort.Ints(list)
		result = list
		return
	}

	pivots := list[:pivotCount]
	sort.Ints(pivots)

	list = list[pivotCount:]
	segments := make([][]int, pivotCount+1)
	found := false

	for _, el := range list {
		found = false
		for pindex, pvalue := range pivots {
			if el < pvalue {
				found = true
				segments[pindex] = append(segments[pindex], el)
				break
			}
		}
		if found == false {
			segments[pivotCount] = append(segments[pivotCount], el)
		}
	}

	var sortedSegment []int
	result = make([]int, n) //preallocate
	i := 0

	for sindex, segment := range segments {
		sortedSegment, err = FivePivot(segment, pivotCount)
		if err != nil {
			return
		}
		i += copy(result[i:], sortedSegment)
		if sindex < len(pivots) {
			result[i] = pivots[sindex]
			i++
		}
	}

	return
}
