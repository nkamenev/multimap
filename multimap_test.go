package multimap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultimap(t *testing.T) {
	m := multimap[string, int]{"key1": {1, 2, 3}, "key2": {4}}

	t.Run("Contains", func(t *testing.T) {
		assert.True(t, m.Contains("key1"))
		assert.False(t, m.Contains("key3"))
	})

	t.Run("Get", func(t *testing.T) {
		items, ok := m.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, []int{1, 2, 3}, items)
		items, ok = m.Get("key3")
		assert.False(t, ok)
		assert.Nil(t, items)
	})

	t.Run("GetAt", func(t *testing.T) {
		item, ok := m.GetAt("key1", 1)
		assert.True(t, ok)
		assert.Equal(t, 2, item)
		item, ok = m.GetAt("key1", 10)
		assert.False(t, ok)
		assert.Equal(t, 0, item)
	})

	t.Run("ForKeys", func(t *testing.T) {
		keys := make(map[string][]int)
		m.ForKeys(func(k string, v []int) { keys[k] = v })
		assert.Equal(t, []int{1, 2, 3}, keys["key1"])
		assert.Equal(t, []int{4}, keys["key2"])
	})

	t.Run("ForKey", func(t *testing.T) {
		var sum int
		m.ForKey("key1", func(v int) { sum += v })
		assert.Equal(t, 6, sum)
	})

	t.Run("For", func(t *testing.T) {
		results := make(map[string]int)
		m.For(func(k string, v int) { results[k] += v })
		assert.Equal(t, 6, results["key1"])
		assert.Equal(t, 4, results["key2"])
	})

	t.Run("NumKeys", func(t *testing.T) {
		assert.Equal(t, 2, m.NumKeys())
	})

	t.Run("IsNil", func(t *testing.T) {
		var nilMap multimap[string, int]
		assert.True(t, nilMap.IsNil())
		assert.False(t, m.IsNil())
	})

	t.Run("Len", func(t *testing.T) {
		assert.Equal(t, 4, m.Len())
	})

	t.Run("LenKey", func(t *testing.T) {
		assert.Equal(t, 3, m.LenKey("key1"))
		assert.Equal(t, 0, m.LenKey("key3"))
	})

	t.Run("Keys", func(t *testing.T) {
		assert.ElementsMatch(t, []string{"key1", "key2"}, m.Keys())
	})
}

func TestMutableMultimap(t *testing.T) {
	m := Make[string, int](2)

	t.Run("SetKey", func(t *testing.T) {
		m.SetKey("key1")
		assert.True(t, m.Contains("key1"))
		items, ok := m.Get("key1")
		assert.True(t, ok)
		assert.Empty(t, items)
	})

	t.Run("Set", func(t *testing.T) {
		m.Set("key1", 1)
		m.Set("key1", 2)
		items, ok := m.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, []int{1, 2}, items)
	})

	t.Run("Delete", func(t *testing.T) {
		m.Set("key2", 3)
		m.Delete("key2")
		assert.False(t, m.Contains("key2"))
	})

	t.Run("DeleteAt", func(t *testing.T) {
		m.Set("key1", 3)
		m.DeleteAt("key1", 1)
		items, ok := m.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, []int{1, 3}, items)
		m.DeleteAt("key1", 10) // Out of bounds
		m.DeleteAt("key3", 0)  // Non-existent key
		items, ok = m.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, []int{1, 3}, items)
	})

	t.Run("Clear", func(t *testing.T) {
		m.Set("key1", 4)
		m.Clear()
		assert.Equal(t, 0, m.NumKeys())
	})

	t.Run("Immutable", func(t *testing.T) {
		m.Set("key1", 5)
		im := m.Immutable()
		items, ok := im.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, []int{5}, items)
	})
}

func TestSequentialMultimap(t *testing.T) {
	m := Make[string, int](2)
	m.Set("key1", 1)
	m.Set("key1", 2)
	m.Set("key1", 3)
	s := m.Sequential()

	t.Run("Start", func(t *testing.T) {
		item, ok := s.Start("key1")
		assert.True(t, ok)
		assert.Equal(t, 1, item)
		item, ok = s.Start("key2")
		assert.False(t, ok)
		assert.Equal(t, 0, item)
	})

	t.Run("End", func(t *testing.T) {
		item, ok := s.End("key1")
		assert.True(t, ok)
		assert.Equal(t, 3, item)
		item, ok = s.End("key2")
		assert.False(t, ok)
		assert.Equal(t, 0, item)
	})
	t.Run("Next", func(t *testing.T) {
		s.Start("key1")
		item, ok := s.Next("key1") // i=0, returns 1, i becomes 1
		assert.True(t, ok)
		assert.Equal(t, 1, item)
		item, ok = s.Next("key1") // i=1, returns 2, i becomes 2
		assert.True(t, ok)
		assert.Equal(t, 2, item)
		item, ok = s.Next("key1") // i=2, returns 3, i becomes 3
		assert.True(t, ok)
		assert.Equal(t, 3, item)
		item, ok = s.Next("key1") // i=3, out of bounds
		assert.False(t, ok)
		assert.Equal(t, 0, item)
	})

	t.Run("Reset", func(t *testing.T) {
		s.Next("key1")             // i=1
		s.Next("key1")             // i=2
		s.Reset("key1")            // i=0
		item, ok := s.Next("key1") // i=0, returns 1
		assert.True(t, ok)
		assert.Equal(t, 1, item)
	})

	t.Run("ResetAll", func(t *testing.T) {
		m.Set("key2", 4)
		s := m.Sequential()
		s.Next("key1")               // i=1 for key1
		s.Next("key2")               // i=1 for key2
		s.ResetAll()                 // i=0 for all keys
		item1, ok1 := s.Next("key1") // i=0, returns 1
		item2, ok2 := s.Next("key2") // i=0, returns 4
		assert.True(t, ok1)
		assert.True(t, ok2)
		assert.Equal(t, 1, item1)
		assert.Equal(t, 4, item2)
	})
}
