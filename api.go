package multimap

type Multimap[K comparable, V any] interface {
	IsNil() bool
	Contains(key K) bool
	Get(key K) ([]V, bool)
	GetAt(key K, i int) (V, bool)
	NumKeys() int
	Len() int
	Keys() []K
	LenKey(key K) int
	ForKey(key K, fn func(V))
	ForKeys(fn func(K, []V))
	For(fn func(K, V))
	Sequential() SequentialMultimap[K, V]
}

type MutableMultimap[K comparable, V any] interface {
	Multimap[K, V]
	SetKey(key K)
	Set(key K, values ...V)
	Delete(key K)
	DeleteAt(key K, i int)
	Clear()
	Immutable() Multimap[K, V]
}

type SequentialMultimap[K comparable, V any] interface {
	Multimap[K, V]
	Start(key K) (V, bool)
	End(key K) (V, bool)
	Next(key K) (V, bool)
	Reset(key K)
	ResetAll()
}
