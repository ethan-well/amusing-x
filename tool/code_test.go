package tool

import (
	"fmt"
	"sort"
	"sync/atomic"
	"testing"
)

func TestSomeCode(t *testing.T) {
	a := []int{1, 9, 5, 8, 19, 10, 3, 13}
	n := maxDepth(10)
	fmt.Println(n)

	sort.Ints(a)
	//sort.Slice()
	//sort.Sort()
	//sort.Float64s()
	//sort.Reverse()
	//sort.SliceStable()
	//sort.IsSorted()
	//sort.Stable()
	//sort.SearchInts()

	var f *[]string
	fmt.Println(f == f, f == nil)

	var value atomic.Value
	//value.Store()
	value.Load()
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		fmt.Println(i) // 10, 5, 2, 1
		depth++
	}
	return depth * 2
}
