package sets

import (
	"sort"
)

type MultiSet map[string]Set

func (s MultiSet) init() MultiSet {
	if s == nil {
		s = MultiSet{}
	}
	return s
}

func (s MultiSet) Add(key string, values ...string) MultiSet {
	s = s.init()
	for _, k := range values {
		s[k] = s[k].Add(values...)
	}
	return s
}

// AddIfNotExists returns true if the key was added, false if it already existed
func (s MultiSet) AddIfNotExists(key, value string) bool {
	if m := s[key]; m == nil {
		return m.AddIfNotExists(value)
	}
	s[key] = SetOf(value)
	return true
}

func (s MultiSet) Clone() MultiSet {
	ns := make(MultiSet, len(s))
	for k, v := range s {
		ns[k] = v.Clone()
	}
	return ns
}

func (s MultiSet) Values(key string) Set {
	return s[key]
}

func (s MultiSet) Merge(o MultiSet) MultiSet {
	s = s.init()
	for k, v := range o {
		s[k] = s[k].Merge(v)
	}
	return s
}

func (s MultiSet) MergeSet(key string, o Set) MultiSet {
	s = s.init()
	s[key] = s[key].Merge(o)
	return s
}

func (s MultiSet) Delete(keys ...string) MultiSet {
	for _, k := range keys {
		delete(s, k)
	}
	return s
}

func (s MultiSet) DeleteValues(key string, values ...string) MultiSet {
	m := s[key]
	for _, v := range values {
		delete(m, v)
	}
	if len(m) == 0 {
		delete(s, key)
	}
	return s
}

func (s MultiSet) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s MultiSet) Len() int {
	return len(s)
}

func (s MultiSet) Keys() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s MultiSet) SortedKeys() []string {
	keys := s.Keys()
	sort.Strings(keys)
	return keys
}
