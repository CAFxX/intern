// Package intern interns strings.
// Interning is best effort only.
// Interned strings may be removed automatically
// at any time without notification.
// All functions may be called concurrently
// with themselves and each other.
package intern

import "sync"

var pool = sync.Pool{
	New: func() interface{} {
		return make(map[string]string)
	},
}

// String returns s, interned.
func String(s string) string {
	m := pool.Get().(map[string]string)
	if c, found := m[s]; found {
		s = c
	} else {
		m[s] = s
	}
	pool.Put(m)
	return s
}

// Strings returns the strings, interned
func Strings(ss ...string) []string {
	m := pool.Get().(map[string]string)
	for i, s := range ss {
		if c, found := m[s]; found {
			ss[i] = c
		} else {
			m[s] = s
		}
	}
	pool.Put(m)
	return ss
}

// Bytes returns b converted to a string, interned.
func Bytes(b []byte) (s string) {
	m := pool.Get().(map[string]string)
	if c, found := m[string(b)]; found {
		s = c
	} else {
		s = string(b)
		m[s] = s
	}
	pool.Put(m)
	return s
}
