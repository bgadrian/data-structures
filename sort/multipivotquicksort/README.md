
## multipivotquicksort  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/bgadrian/data-structures/sort/multipivotquicksort)
MultiPivot uses a variant of the QuickSort algorithm with multiple pivots, splitting the arrays in multiple segments (pivots+1). It can be used to sort large arrays.

Dual-pivot quick search is used in Java7 default sorting method. This algorithm is NOT the sme with the multi-way quicksort.
The current implementation works only with int slices, but can be easily modified to compare your own data structures, but be careful at the memory usage.

### How to use
See the **tests, benchmarks and examples** from the package.
Make your own benchmarks with your own data to find the best values for pivotCount and singleThread.

### Papers
**The package was built for academic purposes**, to learn more about newest research made see the following papers:

> http://epubs.siam.org/doi/pdf/10.1137/1.9781611973198.6

 https://web.archive.org/web/20151002230717/
 http://iaroslavski.narod.ru/quicksort/DualPivotQuicksort.pdf
 http://cs.stanford.edu/~rishig/courses/ref/l11a.pdf 
 
 http://rerun.me/2013/06/13/quicksorting-3-way-and-dual-pivot/  
 https://lusy.fri.uni-lj.si/sites/lusy.fri.uni-lj.si/files/publications/alopez2014-seminar-qsort.pdf
 http://eiche.theoinf.tu-ilmenau.de/quicksort-experiments/uploads/multi-pivot-paper-draft-2015-10-12.pdf
 http://iopscience.iop.org/article/10.1088/1757-899X/180/1/012051/pdf

> "First, we have confirmed previous experimental results
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

## Example
```go
    package main
    import "github.com/bgadrian/data-structures/sort/multipivotquicksort"
    import "fmt"
    
    func main(){
        a := []int{5, 3, 4, 7, 1, 9, 3, 4, 9, 1, 4, 7}
        var pivotCount uint8 = 3
        singleThread := false
        res, err := multipivotquicksort.MultiPivot(a, pivotCount, singleThread)
        if err != nil {
            fmt.Errorf("%v", err)
            return
        }
    
        fmt.Printf("Result is\n%v", res)
        //Output: Result is
        //[1 1 3 3 4 4 4 5 7 7 9 9]
	}
```
