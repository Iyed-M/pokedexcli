package pokecache

import (
	"testing"
)

func TestCache_Add(t *testing.T) {
	c := NewCache(40)
	type test struct {
		key string
		val []byte
	}
	tests := []test{
		{"a", []byte("A")},
		{"b", []byte("B")},
		{"c", []byte("C")},
		{"d", []byte("D")},
	}

	for _, tt := range tests {
		c.Add(tt.key, tt.val)
		v, _ := c.Get(tt.key)
		if string(tt.val) != string(v) {
			t.Errorf("error expected %s got %s  ", tt.val, v)
		}
	}
}
