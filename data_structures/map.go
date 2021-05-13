package structures

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// map[string]int
// TODO:
// 1. handle collisions - done
// 2. add resize feature - done
// 3. handle non-ASCII keys - done
type Map struct {
	count    int
	buckets  []*bucket
	resizing bool
}

type bucket struct {
	key  string
	val  int
	next *bucket
}

const (
	initial_size = 8
)

var (
	table [256]uint8 // 256
)

// TODO: comments
func PearsonHash(key string, bucketsSize int) int {
	if key == "" {
		return 0
	}

	var times int
	if bucketsSize <= 256 {
		times = 1
	} else {
		bs := bucketsSize
		for bs != 0 {
			bs = bs / 256
			times++
		}
	}

	var keyAsBytes = []byte(key)
	var firstCharAsInt = int(keyAsBytes[0])

	var finalHash uint64
	for i := 0; i < times; i++ {
		hash := uint8((firstCharAsInt + i) % 256)
		for _, b := range keyAsBytes {
			hash = table[hash^b]
		}
		finalHash = (finalHash << 8) | uint64(hash)
	}
	// possible bug: resulting value may be larger than int (--> overflow, negative idx)
	return int(finalHash % uint64(bucketsSize))
}

func init() {
	rand.Seed(time.Now().UnixNano())
	for idx, v := range rand.Perm(256) {
		table[idx] = uint8(v)
	}
}

func NewMap() *Map {
	return &Map{
		count:   0,
		buckets: make([]*bucket, initial_size),
	}
}

func (m *Map) Insert(key string, val int) {
	if !m.resizing {
		m.checkAndResize()
	}

	idx := PearsonHash(key, len(m.buckets))

	if b := m.buckets[idx]; b != nil {
		for {
			if b.key == key {
				b.val = val
				return
			}
			if b.next == nil {
				m.count++
				b.next = &bucket{
					key: key,
					val: val,
				}
				return
			}
			b = b.next
		}
	} else {
		m.count++
		m.buckets[idx] = &bucket{
			key: key,
			val: val,
		}
	}
}

func (m *Map) Find(key string) (int, bool) {
	idx := PearsonHash(key, len(m.buckets))
	b := m.buckets[idx]

	if b != nil {
		for {
			if b.key == key {
				return b.val, true
			}

			if b.next == nil {
				break
			}
			b = b.next
		}
	}

	return 0, false
}

func (m *Map) Remove(key string) {
	if !m.resizing {
		m.checkAndResize()
	}

	idx := PearsonHash(key, len(m.buckets))
	b := m.buckets[idx]

	if b != nil {
		if b.key == key {
			m.buckets[idx] = b.next
			m.count--
			return
		}

		var prev *bucket
		for {
			if b.key == key {
				prev.next = b.next
				m.count--
				return
			}

			if b.next == nil {
				break
			}
			prev = b
			b = b.next
		}
	}
	return
}

func (m *Map) checkAndResize() {
	size := len(m.buckets)
	loadPercentage := m.count * 100 / size

	if loadPercentage > 70 {
		size = size << 1
		m.resizing = true
	} else if loadPercentage < 10 && size > initial_size {
		size = size >> 1
		m.resizing = true
	}

	if m.resizing {
		m.resize(size)
	}
}

func (m *Map) resize(newSize int) {
	oldBuckets := m.buckets
	m.buckets = make([]*bucket, newSize)
	count := m.count

	for _, bucket := range oldBuckets {
		if bucket == nil {
			continue
		}
		for {
			m.Insert(bucket.key, bucket.val)
			if bucket.next == nil {
				break
			}
			bucket = bucket.next
		}
	}
	m.count = count
	m.resizing = false
}

func (m *Map) str() string {
	rows := make([]string, len(m.buckets))

	for idx, bucket := range m.buckets {
		var rs []string
		if bucket == nil {
			rs = append(rs, "nil")
		} else {
			for {
				rs = append(rs, fmt.Sprintf("{key: '%s', val: %d}", bucket.key, bucket.val))
				if bucket.next == nil {
					break
				}
				bucket = bucket.next
			}
		}
		rows[idx] = fmt.Sprintf("%10d: [%s]", idx, strings.Join(rs, "->"))
	}

	return strings.Join(rows, "\n")
}
