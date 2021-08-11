package sets

import "testing"

func TestMatch(t *testing.T) {
	ms := MultiSet{
		"a": SetOf("1", "2", "3"),
		"b": SetOf("1", "5", "3"),
		"c": SetOf("1"),
	}
	fn := func(key string, s Set) bool {
		return s.Has("2")
	}

	if ms.Match(fn, true) {
		t.Fatal("expected match all to fail")
	}

	if !ms.Match(fn, false) {
		t.Fatal("expected match any to work")
	}
}
