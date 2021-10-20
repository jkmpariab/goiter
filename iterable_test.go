package goiter_test

import (
	"strconv"
	"testing"

	"github.com/jkmpariab/goiter"
)

// //
// // Vector is defined already
// //
// type Vector []int
// type VectorIter struct {
// 	vector *Vector
// 	index  int
// }
//
// func (v *Vector) Iter() goiter.Iterator {
// 	return &VectorIter{v, 0}
// }
//
// func (vi *VectorIter) Next() (next interface{}, has bool) {
// 	if vi.index < len(*vi.vector) {
// 		next = ([]int)(*vi.vector)[vi.index]
// 		has = true
// 	}
// 	vi.index++
// 	return
// }

func TestIterableMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedNums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.Map(func(v interface{}) interface{} {
		n := v.(int)
		n *= 2
		return n
	}).ForEach(func(v interface{}) {
		if expectedNums[i] != v.(int) {
			t.Errorf("expected: %d, got: %d\n", expectedNums[i], v.(int))
		}

		i++
	})
}

func TestIterableMapString(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedStrings := []string{"2", "4", "6", "8", "10", "12", "14", "16", "18"}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.Map(func(v interface{}) interface{} {
		n := v.(int)
		n *= 2
		return strconv.Itoa(n)
	}).ForEach(func(v interface{}) {
		if expectedStrings[i] != v.(string) {
			t.Errorf("expected: %q, got: %q\n", expectedStrings[i], v.(string))
		}

		i++
	})
}

func TestIterableFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedNums := []int{1, 3, 5, 7, 9}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.Filter(func(v interface{}) bool {
		n := v.(int)
		return n%2 != 0 // filter odd numbers
	}).ForEach(func(v interface{}) {
		if expectedNums[i] != v.(int) {
			t.Errorf("expected: %d, got: %d\n", expectedNums[i], v.(int))
		}

		i++
	})
}

func TestIterableFilterAndMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedNums := []int{2, 6, 10, 14, 18}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.Filter(func(v interface{}) bool {
		n := v.(int)
		return n%2 != 0
	}).Map(func(v interface{}) interface{} {
		n := v.(int)
		n *= 2
		return n
	}).ForEach(func(v interface{}) {
		if expectedNums[i] != v.(int) {
			t.Errorf("expected: %d, got: %d\n", expectedNums[i], v.(int))
		}

		i++
	})
}

func TestIterableFilterAndMapString(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedStrings := []string{"2", "6", "10", "14", "18"}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.Filter(func(v interface{}) bool {
		n := v.(int)
		return n%2 != 0
	}).Map(func(v interface{}) interface{} {
		n := v.(int)
		n *= 2
		return strconv.Itoa(n)
	}).ForEach(func(v interface{}) {
		if expectedStrings[i] != v.(string) {
			t.Errorf("expected: %q, got: %q\n", expectedStrings[i], v.(string))
		}

		i++
	})
}

func TestIterableWhile(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedNums := []int{1, 2, 3, 4, 5}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.While(func(v interface{}) bool {
		n := v.(int)
		return n <= 5
	}).ForEach(func(v interface{}) {
		if expectedNums[i] != v.(int) {
			t.Errorf("expected: %d, got: %d\n", expectedNums[i], v.(int))
		}

		i++
	})
}

func TestIterableSkip(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedNums := []int{6, 7, 8, 9}

	iterator := SliceIterator(nums)
	iterable := goiter.NewIterable(iterator)

	i := 0
	iterable.Skip(5).ForEach(func(v interface{}) {
		if expectedNums[i] != v.(int) {
			t.Errorf("expected: %d, got: %d\n", expectedNums[i], v.(int))
		}

		i++
	})
}
