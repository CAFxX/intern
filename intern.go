// Package intern interns strings.
// Interning is best effort only.
// Interned strings may be removed automatically
// at any time without notification.
// All functions may be called concurrently
// with themselves and each other.
package intern

import "sync"

type interningTable map[string]string

var pool sync.Pool = sync.Pool{
	New: func() interface{} {
		return make(interningTable)
	},
}

// String returns s, interned.
func String(s string) string {
	m := pool.Get().(interningTable)
	is := m.intern(s)
	pool.Put(m)
	return is
}

// Strings returns, for each string in the provided slice, the corresponding interned string
func Strings(ss ...string) []string {
	m := pool.Get().(interningTable)
	for i, s := range ss {
		ss[i] = m.intern(s)
	}
	pool.Put(m)
	return ss
}

func (m interningTable) intern(s string) string {
	if c, ok := m[s]; ok {
		return c
	}
	m[s] = s
	return s
}

// Bytes returns b converted to a string, interned.
func Bytes(b []byte) string {
	return String(string(b))
}
