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

func TestForEach(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := Vector(nums)
	iterator := v.Iter()

	i := 0
	iter.ForEach(iterator, func(v interface{}) {
		if v.(int) != nums[i] {
			t.Errorf("expected: %d, got: %d\n", nums[i], v.(int))
		}

		i++
	})
}

func TestCollect(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := Vector(nums)
	iterator := v.Iter()

	for i, v := range iter.Collect(iterator) {
		if v.(int) != nums[i] {
			t.Errorf("expected: %d, got: %d\n", nums[i], v.(int))
		}
	}
}

func TestAny(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := Vector(nums)
	iterator := v.Iter()

	any := iter.Any(iterator, func(v interface{}) bool {
		n := v.(int)
		return n == 5
	})

	if !any {
		t.Errorf("expected any '5' in slice []int{1, 2, 3, 4, 5, 6, 7, 8, 9}\n")
	}
}

func TestAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := Vector(nums)
	iterator := v.Iter()

	all := iter.All(iterator, func(v interface{}) bool {
		n := v.(int)
		return n > 0 && n < 10
	})

	if !all {
		t.Errorf("expected all values in slice []int{1, 2, 3, 4, 5, 6, 7, 8, 9} are greater than '0' and less than '10'\n")
	}
}
