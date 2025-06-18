package multimap

// Multimap maps keys to multiple values.
type multimap[K comparable, V any] map[K][]V

// Contains checks if the key exists.
func (m multimap[K, V]) Contains(key K) bool {
	_, contains := m[key]
	return contains
}

// Get returns the values for a key, or nil if absent.
func (m multimap[K, V]) Get(key K) ([]V, bool) {
	items, exists := m[key]
	if !exists {
		return nil, false
	}
	return items, true
}

// GetAt returns the value at index i for a key, or zero if invalid.
func (m multimap[K, V]) GetAt(key K, i int) (V, bool) {
	items, exists := m[key]
	if !exists {
		var zero V
		return zero, false
	}
	if i >= len(items) {
		var zero V
		return zero, false
	}
	return items[i], true
}

// ForKeys iterates over all keys and their values.
func (m multimap[K, V]) ForKeys(fn func(K, []V)) {
	for k, items := range m {
		fn(k, items)
	}
}

// ForKey iterates over values for a specific key.
func (m multimap[K, V]) ForKey(key K, fn func(V)) {
	items, ok := m[key]
	if !ok {
		return
	}
	for i := 0; i < len(items); i++ {
		fn(items[i])
	}
}

// For iterates over all key-value pairs.
func (m multimap[K, V]) For(fn func(K, V)) {
	for k, items := range m {
		for _, item := range items {
			fn(k, item)
		}
	}
}

// NumKeys returns the number of keys.
func (m multimap[K, V]) NumKeys() int {
	return len(m)
}

// IsNil checks if the multimap is nil.
func (m multimap[K, V]) IsNil() bool {
	return m == nil
}

// Len returns the total number of values across all keys.
func (m multimap[K, V]) Len() int {
	var sum int
	for _, items := range m {
		sum += len(items)
	}
	return sum
}

// LenKey returns the number of values for a key.
func (m multimap[K, V]) LenKey(key K) int {
	return len(m[key])
}

// Keys returns a slice of all keys.
func (m multimap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Sequential wraps the multimap for sequential access.
func (m multimap[K, V]) Sequential() SequentialMultimap[K, V] {
	return makeSequential(m)
}
