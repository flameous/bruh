package structures

import (
	"strconv"
	"testing"
)

func TestPearsonHashing(t *testing.T) {

	for _, mapSize := range []int{8, 16, 32, 1024, 2048, 4096, 131072} {

		for _, key := range []string{"", "foo", "bar", "baz", "qux"} {
			idx := PearsonHash(key, mapSize)
			if idx >= mapSize {
				t.Errorf("hash func is broken: idx (%d) >= size (%d)", idx, mapSize)
			}
		}
	}
}

func TestMap_InsertAndFind(t *testing.T) {
	m := NewMap()

	m.Insert("foo", 1)

	val, ok := m.Find("foo")
	if !ok {
		t.Error("value not found")
	}
	if val != 1 {
		t.Errorf("expected value = 1, but found %d (because collision and overriding the data)", val)
	}

	_, ok = m.Find("bar")
	if ok {
		t.Error("expected false for non-existent key")
	}
}

func TestMap_Collisions_and_Resizing(t *testing.T) {
	m := NewMap()

	m.Insert("foo", 1)
	m.Insert("bar", -2)
	m.Insert("qux", 2)
	m.Insert("quux", 3)
	m.Insert("quuux", 4)
	m.Insert("quuuux", 5)
	m.Insert("quuuuux", 6)
	m.Insert("xyz", 7)

	val, ok := m.Find("foo")
	if !ok {
		t.Error("value not found")
	}
	if val != 1 {
		t.Errorf("expected value = 1, but found %d (because collision and overriding the data)", val)
	}

	if m.count != 8 {
		t.Errorf("there must be 8 elements, but counter shows %d", m.count)
	}

	if len(m.buckets) != 16 {
		t.Errorf("buckets size expected to be 16")
	}
}

func TestMap_NonASCIIcharacters(t *testing.T) {
	m := NewMap()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("caught panic: %v", r)
		}
	}()
	m.Insert("フルーツバスケット", 1)

	val, ok := m.Find("フルーツバスケット")
	if !ok {
		t.Errorf("value not found")
	}

	if val != 1 {
		t.Errorf("invalid value")
	}
}

func TestMap_Remove(t *testing.T) {
	m := NewMap()

	m.Insert("foo", 1)
	m.Insert("bar", -2)
	m.Insert("qux", 2)
	m.Insert("abc", 3)
	m.Insert("def", 4)
	m.Insert("xyz", 5)

	_, ok := m.Find("xyz")
	if !ok {
		t.Error("value not found")
	}

	m.Remove("xyz")

	_, ok = m.Find("xyz")
	if ok {
		t.Error("value found but expected to be deleted")
	}
}

func TestMap_functional(t *testing.T) {
	builtinMap := make(map[string]int)
	m := NewMap()

	for i := 0; i < 100; i++ {
		key := strconv.Itoa(i) + "_suffix"

		builtinMap[key] = i
		m.Insert(key, i)

		if len(builtinMap) != m.count {
			t.Error("different sizes of builtin map and implemented one")
		}
	}

	for key, val := range builtinMap {
		v, ok := m.Find(key)
		if !ok {
			t.Errorf("key '%s' not found in implemented map", key)
		}
		if val != v {
			t.Errorf("data corrupted. expected: %d, but got: %d", val, v)
		}
	}

	for key := range builtinMap {
		m.Remove(key)

		val, ok := m.Find(key)
		if ok {
			t.Errorf("key '%s' val: %d expected to be deleted", key, val)
		}
	}

	if m.count != 0 {
		t.Errorf("impl. map expected to be empty but it contains:\n%s\ncount = %d", m.str(), m.count)
	}
}
