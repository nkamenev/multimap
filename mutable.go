package multimap

// MutableMultimap is a mutable multimap embedding a multimap.
type mutMultimap[K comparable, V any] struct{ multimap[K, V] }

// Make creates a new MutableMultimap with initial capacity.
func Make[K comparable, V any](size int) MutableMultimap[K, V] {
	return mutMultimap[K, V]{make(multimap[K, V], size)}
}

// SetKey initializes an empty value list for a key.
func (m mutMultimap[K, V]) SetKey(key K) {
	m.multimap[key] = []V{}
}

// Set appends a value to a key's value list.
func (m mutMultimap[K, V]) Set(key K, values ...V) {
	items, ok := m.multimap[key]
	if !ok {
		m.multimap[key] = values
		return
	}
	m.multimap[key] = append(items, values...)
}

// Delete removes a key and its values.
func (m mutMultimap[K, V]) Delete(key K) {
	delete(m.multimap, key)
}

// DeleteAt removes a value at index i for a key.
func (m mutMultimap[K, V]) DeleteAt(key K, i int) {
	items, ok := m.multimap[key]
	if !ok {
		return
	}
	if i >= len(items) {
		return
	}
	m.multimap[key] = append(items[:i], items[i+1:]...)
}

// Clear removes all keys and values.
func (m mutMultimap[K, V]) Clear() {
	clear(m.multimap)
}

// Immutable returns an immutable view of the multimap.
func (m mutMultimap[K, V]) Immutable() Multimap[K, V] {
	return m.multimap
}
