package iter

type Iterator interface {
	Next() bool
	Value() interface{}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type mapIter struct {
	iterator Iterator
	fn       func(interface{}) interface{}
	nv       interface{}
}

func (mi *mapIter) Next() bool {
	if mi.iterator.Next() {
		mi.nv = mi.fn(mi.iterator.Value())
		return true
	}
	return false
}

func (mi *mapIter) Value() interface{} {
	return mi.nv
}

type filterIter struct {
	iterator Iterator
	fn       func(interface{}) bool
	nv       interface{}
}

func (fi *filterIter) Next() bool {
	if fi.iterator.Next() {
		fi.nv = fi.iterator.Value()
		if fi.fn(fi.nv) {
			return true
		} else {
			return fi.Next()
		}
	}
	return false
}

func (fi *filterIter) Value() interface{} {
	return fi.nv
}

type whileIter struct {
	iterator Iterator
	fn       func(interface{}) bool
	nv       interface{}
}

func (wi *whileIter) Next() bool {
	if wi.iterator.Next() {
		wi.nv = wi.iterator.Value()
		if wi.fn(wi.nv) {
			return true
		}
	}
	return false
}

func (wi *whileIter) Value() interface{} {
	return wi.nv
}

type skipIter struct {
	iterator Iterator
	skip     uint
	count    uint
	nv       interface{}
}

func (si *skipIter) Next() bool {
	if hasNext := si.iterator.Next(); hasNext {
		si.nv = si.iterator.Value()
		if si.count < si.skip {
			si.count++
			return si.Next()
		}
		si.count++
		return hasNext
	}
	return false
}

func (si *skipIter) Value() interface{} {
	return si.nv
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

// Consumer
func ForEach(iterator Iterator, f func(interface{})) {
	for iterator.Next() {
		f(iterator.Value())
	}
}

// Consumer
func Collect(iterator Iterator) []interface{} {
	var collected []interface{}
	for iterator.Next() {
		collected = append(collected, iterator.Value())
	}
	return collected
}

// Consumer
func Any(iterator Iterator, f func(interface{}) bool) bool {
	for iterator.Next() {
		if f(iterator.Value()) {
			return true
		}
	}
	return false
}

// Consumer
func All(iterator Iterator, f func(interface{}) bool) bool {
	for iterator.Next() {
		if !f(iterator.Value()) {
			return false
		}
	}
	return true
}

// Adaptor
func Map(iterator Iterator, f func(interface{}) interface{}) Iterator {
	return &mapIter{iterator: iterator, fn: f}
}

// Adaptor
func Filter(iterator Iterator, f func(interface{}) bool) Iterator {
	return &filterIter{iterator: iterator, fn: f}
}

// Adaptor
func While(iterator Iterator, f func(interface{}) bool) Iterator {
	return &whileIter{iterator: iterator, fn: f}
}

// Adaptor
func Skip(iterator Iterator, skip uint) Iterator {
	return &skipIter{iterator: iterator, skip: skip}
}
