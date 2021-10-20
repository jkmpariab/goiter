package goiter_test

import (
	"testing"

	"github.com/jkmpariab/goiter"
)

func SliceIterator(slc []int) *sliceIter {
	return &sliceIter{slc: slc}
}

type sliceIter struct {
	slc   []int
	index int
	nv    int
}

func (si *sliceIter) Next() (hasNext bool) {
	if si.index < len(si.slc) {
		si.nv = si.slc[si.index]
		hasNext = true
	}
	si.index++
	return
}

func (vi *sliceIter) Value() interface{} {
	return vi.nv
}

func TestForEach(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iterator := SliceIterator(nums)

	i := 0
	goiter.ForEach(iterator, func(v interface{}) {
		if v.(int) != nums[i] {
			t.Errorf("expected: %d, got: %d\n", nums[i], v.(int))
		}

		i++
	})
}

func TestCollect(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iterator := SliceIterator(nums)

	for i, v := range goiter.Collect(iterator) {
		if v.(int) != nums[i] {
			t.Errorf("expected: %d, got: %d\n", nums[i], v.(int))
		}
	}
}

func TestAny(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iterator := SliceIterator(nums)

	any := goiter.Any(iterator, func(v interface{}) bool {
		n := v.(int)
		return n == 5
	})

	if !any {
		t.Errorf("expected any '5' in slice []int{1, 2, 3, 4, 5, 6, 7, 8, 9}\n")
	}
}

func TestAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iterator := SliceIterator(nums)

	all := goiter.All(iterator, func(v interface{}) bool {
		n := v.(int)
		return n > 0 && n < 10
	})

	if !all {
		t.Errorf("expected all values in slice []int{1, 2, 3, 4, 5, 6, 7, 8, 9} are greater than '0' and less than '10'\n")
	}
}
