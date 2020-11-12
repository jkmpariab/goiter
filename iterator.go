package iter

type Iterator interface {
	Next() (next interface{}, has bool)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type mapIterator struct {
	iterator Iterator
	mapf     func(interface{}) interface{}
}

func (mi *mapIterator) Next() (next interface{}, has bool) {
	if next, has := mi.iterator.Next(); has {
		return mi.mapf(next), true
	}
	return nil, false
}

type filterIterator struct {
	iterator Iterator
	filter   func(interface{}) bool
}

func (fi *filterIterator) Next() (next interface{}, has bool) {
	if next, has := fi.iterator.Next(); has {
		if fi.filter(next) {
			return next, true
		} else {
			return fi.Next()
		}
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

// Adaptor
func Map(iterator Iterator, f func(interface{}) interface{}) Iterator {
	return &mapIterator{iterator, f}
}

// Adaptor
func Filter(iterator Iterator, f func(interface{}) bool) Iterator {
	return &filterIterator{iterator, f}
}
