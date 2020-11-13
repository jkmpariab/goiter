package iter

type Iterator interface {
	Next() (next interface{}, has bool)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type mapIterator struct {
	iterator Iterator
	mapf     func(interface{}) interface{}
}

func (mi *mapIterator) Next() (interface{}, bool) {
	if next, has := mi.iterator.Next(); has {
		return mi.mapf(next), true
	}
	return nil, false
}

type filterIterator struct {
	iterator Iterator
	filter   func(interface{}) bool
}

func (fi *filterIterator) Next() (interface{}, bool) {
	if next, has := fi.iterator.Next(); has {
		if fi.filter(next) {
			return next, true
		} else {
			return fi.Next()
		}
	}
	return nil, false
}

type whileIterator struct {
	iterator Iterator
	while    func(interface{}) bool
}

func (wi *whileIterator) Next() (interface{}, bool) {
	if next, has := wi.iterator.Next(); has {
		if wi.while(next) {
			return next, true
		}
	}
	return nil, false
}

type skipIterator struct {
	iterator Iterator
	skip     uint
	count    uint
}

func (si *skipIterator) Next() (interface{}, bool) {
	if next, has := si.iterator.Next(); has {
		if si.count < si.skip {
			si.count++
			return si.Next()
		}
		si.count++
		return next, has
	}
	return nil, false
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

// Consumer
func ForEach(iterator Iterator, f func(interface{})) {
	for next, has := iterator.Next(); has; next, has = iterator.Next() {
		f(next)
	}
}

// Consumer
func Collect(iterator Iterator) []interface{} {
	var collected []interface{}
	for next, has := iterator.Next(); has; next, has = iterator.Next() {
		collected = append(collected, next)
	}
	return collected
}

// Consumer
func Any(iterator Iterator, f func(interface{}) bool) bool {
	for next, has := iterator.Next(); has; next, has = iterator.Next() {
		if f(next) {
			return true
		}
	}
	return false
}

// Consumer
func All(iterator Iterator, f func(interface{}) bool) bool {
	for next, has := iterator.Next(); has; next, has = iterator.Next() {
		if !f(next) {
			return false
		}
	}
	return true
}

// Adaptor
func Map(iterator Iterator, f func(interface{}) interface{}) Iterator {
	return &mapIterator{iterator, f}
}

// Adaptor
func Filter(iterator Iterator, f func(interface{}) bool) Iterator {
	return &filterIterator{iterator, f}
}

// Adaptor
func While(iterator Iterator, f func(interface{}) bool) Iterator {
	return &whileIterator{iterator, f}
}

// Adaptor
func Skip(iterator Iterator, skip uint) Iterator {
	return &skipIterator{iterator: iterator, skip: skip}
}
