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

func (it *Iterable) Map(f func(interface{}) interface{}) *Iterable {
	return &Iterable{Map(it.iterator, f)}
}

func (it *Iterable) Filter(f func(interface{}) bool) *Iterable {
	return &Iterable{Filter(it.iterator, f)}
}
