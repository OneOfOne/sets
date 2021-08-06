package sets

import (
	"encoding/json"
	"sort"
	"sync"
)

type SafeSet struct {
	s   Set
	mux sync.RWMutex
}

func (ss *SafeSet) Set(keys ...string) {
	ss.mux.Lock()
	ss.s.Set(keys...)
	ss.mux.Unlock()
}

func (ss *SafeSet) MergeSafe(o *SafeSet) {
	ss.mux.Lock()
	o.mux.Lock()
	ss.s.Merge(o.s)
	o.mux.Unlock()
	ss.mux.Unlock()
}

func (ss *SafeSet) Merge(o Set) {
	ss.mux.Lock()
	ss.s.Merge(o)
	ss.mux.Unlock()
}

func (ss *SafeSet) Delete(keys ...string) {
	ss.mux.Lock()
	ss.s.Delete(keys...)
	ss.mux.Unlock()
}

func (ss *SafeSet) Has(key string) bool {
	ss.mux.RLock()
	ok := ss.s.Has(key)
	ss.mux.RUnlock()
	return ok
}

func (ss *SafeSet) AddIfNotExists(key string) bool {
	ss.mux.Lock()
	added := ss.s.AddIfNotExists(key)
	ss.mux.Unlock()
	return added
}

func (ss *SafeSet) Len() int {
	ss.mux.RLock()
	ln := ss.s.Len()
	ss.mux.RUnlock()
	return ln
}

func (ss *SafeSet) Keys() []string {
	ss.mux.RLock()
	keys := ss.s.Keys()
	ss.mux.RUnlock()
	return keys
}

func (ss *SafeSet) SortedKeys() []string {
	keys := ss.Keys()
	sort.Strings(keys)
	return keys
}

func (ss *SafeSet) MarshalJSON() ([]byte, error) {
	keys := ss.Keys()
	return json.Marshal(keys)
}

func (ss *SafeSet) UnmarshalJSON(data []byte) error {
	var s Set
	if err := s.UnmarshalJSON(data); err != nil {
		return err
	}
	ss.mux.Lock()
	ss.s = s
	ss.mux.Unlock()
	return nil
}
