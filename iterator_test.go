package iter_test

import (
	"testing"

	"github.com/jkmpariab/iter"
)

type Vector []int
type VectorIter struct {
	vector *Vector
	index  int
}

func (v *Vector) Iter() iter.Iterator {
	return &VectorIter{v, 0}
}

func (vi *VectorIter) Next() (next interface{}, has bool) {
	if vi.index < len(*vi.vector) {
		next = ([]int)(*vi.vector)[vi.index]
		has = true
	}
	vi.index++
	return
}

func TestIterator(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := Vector(nums)
	iterator := v.Iter()

	m := iter.Map(iterator, func(v interface{}) interface{} {
		item := v.(int)
		item *= 2
		return item
	})

	for i, v := range iter.Collect(m) {
		if nums[i]*2 != v.(int) {
			t.Errorf("expected: %d*2, got: %d\n", nums[i], v.(int))
		}
	}
}

func TestIteratorFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{6, 7, 8, 9}

	v := Vector(nums)

	iterator := v.Iter() // VectorIter{ &v, 0 }

	filterIterator := iter.Filter(iterator, func(v interface{}) bool { return v.(int) > 5 }) // filterIterator{ iterator, func(..) bool {...} }

	// ForEach(filterIterator, func(v interface{}) { fmt.Printf("[TestIteratorFilter] %v\n", v) }) // consume iterator
	for i, v := range iter.Collect(filterIterator) {
		if v.(int) != expected[i] {
			t.Errorf("expected: %d, got: %d\n", expected[i], v.(int))
		}
	}
}
