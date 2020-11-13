package iter

type Iterable struct {
	iterator Iterator
}

func NewIterable(iterator Iterator) *Iterable {
	return &Iterable{iterator}
}

func (it *Iterable) ForEach(f func(interface{})) {
	ForEach(it.iterator, f)
}

func (it *Iterable) Collect() []interface{} {
	return Collect(it.iterator)
}

func (it *Iterable) Any(f func(interface{}) bool) bool {
	return Any(it.iterator, f)
}

func (it *Iterable) All(f func(interface{}) bool) bool {
	return All(it.iterator, f)
}

func (it *Iterable) Map(f func(interface{}) interface{}) *Iterable {
	return &Iterable{Map(it.iterator, f)}
}

func (it *Iterable) Filter(f func(interface{}) bool) *Iterable {
	return &Iterable{Filter(it.iterator, f)}
}

func (it *Iterable) While(f func(interface{}) bool) *Iterable {
	return &Iterable{While(it.iterator, f)}
}

func (it *Iterable) Skip(skip uint) *Iterable {
	return &Iterable{Skip(it.iterator, skip)}
}
