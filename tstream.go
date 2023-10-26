package gocombinatorics

// TStream is like Stream but it emits tuples of type T. The zero value
// emits no tuples. Copying a TStream is not supported and may lead to
// errors.
type TStream[T any] struct {
	items   []T
	indexes []int
	stream  Stream
}

// TCombinations yields all the ways you can pick k items from the items
// slice without replacement where order does not matter.
func TCombinations[T any](items []T, k int) *TStream[T] {
	return newTStream(items, k, Combinations)
}

// TCombinationsWithReplacement yields all the ways you can pick k items
// from the items slice with replacement where order does not matter.
func TCombinationsWithReplacement[T any](items []T, k int) *TStream[T] {
	return newTStream(items, k, CombinationsWithReplacement)
}

// TPermutations yields all the ways you can pick k items from the items
// slice without replacement where order matters.
func TPermutations[T any](items []T, k int) *TStream[T] {
	return newTStream(items, k, Permutations)
}

// TProduct yields all the ways you can pick k items from the items
// slice with replacement where order matters.
func TProduct[T any](items []T, k int) *TStream[T] {
	return newTStream(items, k, Product)
}

func newTStream[T any](
	items []T, k int, streamType func(n, k int) Stream) *TStream[T] {
	stream := streamType(len(items), k)
	indexes := make([]int, stream.TupleSize())
	return &TStream[T]{
		items:   append([]T(nil), items...),
		indexes: indexes,
		stream:  stream,
	}
}

// Next populates values with the next tuple and returns true. If there are
// no more tuples, Next returns false and leaves values unchanged. Caller
// must pass in a slice big enough to hold a tuple.
func (t *TStream[T]) Next(values []T) bool {
	if len(values) < len(t.indexes) {
		panic(kSliceTooSmall)
	}
	if t.stream == nil || !t.stream.Next(t.indexes) {
		return false
	}
	for i := range t.indexes {
		values[i] = t.items[t.indexes[i]]
	}
	return true
}

// TupleSize returns the size of tuples this TStream emits. Caller must
// pass a slice of at least this size to the Next method.
func (t *TStream[T]) TupleSize() int {
	return len(t.indexes)
}

// Reset resets this TStream to the state it had when it was first
// created. After calling Reset, Next will yield the first tuple.
func (t *TStream[T]) Reset() {
	if t.stream != nil {
		t.stream.Reset()
	}
}
