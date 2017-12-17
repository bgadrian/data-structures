/*Package multipivotquicksort
MultiPivot uses a variant of the QuickSort with multiple pivots, splitting the arrays in multiple segments (pivots+1).

Dual-pivot quick search is used in Java7 default sorting method. This algorithm is NOT the sme with the multi-way quicksort.
The current implementation works only with int slices, but can be easily modified to compare your own data structures, but be careful at the memory usage.

The package was built for academic purposes, to learn more about newest research made see the following papers:

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
	"sync"
)

/*MultiPivot uses a variant of the QuickSort with multiple pivots, splitting the arrays in multiple segments (pivots+1).
It consumes more space (memory), is not yet optimized to work with only 1 slice, it copies the data in each step.
singleThread should be used for small data sets, benchmark with your own data sets to see which is best.
pivotCount 1 has the same effect as the regular quickSort, with larger data sets (millions) more then 7 pivots can improve the algorithm.
.*/
func MultiPivot(list []int, pivotCount uint8, singleThread bool) (result []int, err error) {
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

	for _, el := range list {
		for pindex, pvalue := range pivots {
			if el < pvalue {
				segments[pindex] = append(segments[pindex], el)
				goto found
			}
		}
		//the element is >= last pivot, so it goes to the last bucket/segment
		segments[pivotCount] = append(segments[pivotCount], el)
	found:
	}

	//apply the same alg to each segment/bucket
	sortedSegments := make([][]int, len(segments))
	if singleThread {
		for sindex, segment := range segments {
			sortedSegments[sindex], err = MultiPivot(segment, pivotCount, singleThread)
			if err != nil {
				return
			}
		}
	} else {
		var wg sync.WaitGroup
		wg.Add(len(segments))
		//goroutines share memory, has a race problem
		//but each one has its own index
		for sindex, segment := range segments {
			go func(sindex int, segment []int) {
				sortedSegments[sindex], err = MultiPivot(segment, pivotCount, singleThread)
				wg.Done()
			}(sindex, segment)
		}

		wg.Wait()
	}

	//glue the segments in the result
	result = make([]int, n) //preallocate
	i := 0
	for sindex, sortedSegment := range sortedSegments {
		i += copy(result[i:], sortedSegment)
		if sindex < len(pivots) {
			result[i] = pivots[sindex]
			i++
		}
	}

	return
}
