package sets

import (
	"encoding/json"
	"sort"
)

var empty struct{}

func SetOf(keys ...string) Set {
	s := Set{}
	s.Add(keys...)
	return s
}

// Set is a simple set.
type Set map[string]struct{}

func (s Set) init() Set {
	if s == nil {
		s = Set{}
	}
	return s
}

func (s *Set) Set(keys ...string) Set {
	ss := s.Add(keys...)
	if *s == nil {
		*s = ss
	}
	return ss
}

func (s Set) Add(keys ...string) Set {
	s = s.init()
	for _, k := range keys {
		s[k] = empty
	}
	return s
}

// AddIfNotExists returns true if the key was added, false if it already existed
func (s *Set) AddIfNotExists(key string) bool {
	sm := s.init()
	if s == nil {
		*s = sm
	}
	if _, ok := sm[key]; ok {
		return false
	}

	sm[key] = empty
	return true
}

func (s Set) Clone() Set {
	ns := make(Set, len(s))
	for k, v := range s {
		ns[k] = v
	}
	return ns
}

func (s Set) Merge(o Set) Set {
	s = s.init()
	for k := range o {
		s[k] = empty
	}
	return s
}

func (s Set) Delete(keys ...string) Set {
	for _, k := range keys {
		delete(s, k)
	}
	return s
}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Keys() []string {
	if s == nil {
		return nil
	}
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set) SortedKeys() []string {
	keys := s.Keys()
	sort.Strings(keys)
	return keys
}

func (s Set) MarshalJSON() ([]byte, error) {
	keys := s.Keys()
	return json.Marshal(keys)
}

func (s *Set) UnmarshalJSON(data []byte) (err error) {
	var keys []string
	if err = json.Unmarshal(data, &keys); err == nil {
		s.Set(keys...)
	}
	return
}
