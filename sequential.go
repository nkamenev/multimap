package multimap

// SequentialMultimap extends multimap with sequential access.
type seqMultimap[K comparable, V any] struct {
	multimap[K, V]
	seq map[K]int
}

// makeSequential creates a SequentialMultimap from a multimap.
func makeSequential[K comparable, V any](m multimap[K, V]) SequentialMultimap[K, V] {
	return seqMultimap[K, V]{
		multimap: m,
		seq:      make(map[K]int, len(m)),
	}
}

// Start sets iterator to the first value for a key and returns it.
func (m seqMultimap[K, V]) Start(key K) (V, bool) {
	items, ok := m.Get(key)
	if !ok {
		var zero V
		return zero, false
	}
	m.seq[key] = 0
	return items[0], true
}

// End sets iterator to the last value for a key and returns it.
func (m seqMultimap[K, V]) End(key K) (V, bool) {
	items, ok := m.Get(key)
	if !ok || len(items) == 0 {
		var zero V
		return zero, false
	}
	m.seq[key] = len(items) - 1
	return items[len(items)-1], true
}

// Next returns the current value and advances the iterator.
func (m seqMultimap[K, V]) Next(key K) (V, bool) {
	items, ok := m.Get(key)
	if !ok {
		var zero V
		return zero, false
	}
	i := m.seq[key]
	if i >= len(items) {
		var zero V
		return zero, false
	}
	m.seq[key]++
	return items[i], true
}

// Reset sets the iterator to the start for a key.
func (m seqMultimap[K, V]) Reset(key K) {
	m.seq[key] = 0
}

// ResetAll sets iterators to the start for all keys.
func (m seqMultimap[K, V]) ResetAll() {
	for k := range m.seq {
		m.seq[k] = 0
	}
}
