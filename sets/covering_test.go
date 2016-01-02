package sets

import "testing"

func newSetFromStringsArray(t *testing.T, values []string, matching func(interface{}, interface{}) bool) Set {
	set, err := NewSet(matching)
	switch {
	case err != nil:
		t.Fatal("Error not nil: ", err)
	case !set.match(2, 2):
		t.Fatal("2 and 2 don't match!")
	case set.match(5, 1):
		t.Fatal("5 and 1 match!")
	case set.Size() != 0:
		t.Fatal("Wrong initial set size: ", set.Size())
	}

	for _, i := range(values) {
		set.Insert(i)
	}
	if set.Size() != len(values) {
		t.Fatal("Wrong filled set size: ", set.Size())
	}
	return set
}

func newSetFromKSetsArray(t *testing.T, values []KSet, matching func(interface{}, interface{}) bool) Set {
	set, err := NewSet(matching)
	switch {
	case err != nil:
		t.Fatal("Error not nil: ", err)
	//case !set.match(*KSet{key: "a"}, *KSet{key: "a"}):
	//	t.Fatal("a and a don't match!")
	//case !set.match(*KSet{key: "a"}, *KSet{key: "b"}):
	//	t.Fatal("a and b match!")
	case set.Size() != 0:
		t.Fatal("Wrong initial set size: ", set.Size())
	}

	for _, i := range(values) {
		set.Insert(i)
	}
	if set.Size() != len(values) {
		t.Fatal("Wrong filled set size: ", set.Size())
	}
	return set
}


func TestCovering(t *testing.T) {
	s := newSetFromStringsArray(
		t,
		[]string {"a","b","c","d","e","f","g","h","i","j","k","l"},
		func(a, b interface{}) bool {return a == b},
	)
	a1 := newSetFromStringsArray(
		t,
		[]string{"a","b","c","d"},
		func(a, b interface{}) bool {return a == b},
	)
	a2 := newSetFromStringsArray(
		t,
		[]string{"e","f","g","h","i"},
		func(a, b interface{}) bool {return a == b},
	)
	a3 := newSetFromStringsArray(
		t,
		[]string{"j","k","l"},
		func(a, b interface{}) bool {return a == b},
	)
	a4 := newSetFromStringsArray(
		t,
		[]string{"a", "e"},
		func(a, b interface{}) bool {return a == b},
	)
	a5 := newSetFromStringsArray(
		t,
		[]string{"b", "f", "g"},
		func(a, b interface{}) bool {return a == b},
	)
	a6 := newSetFromStringsArray(
		t,
		[]string{"c", "d", "g", "h", "k", "l"},
		func(a, b interface{}) bool {return a == b},
	)
	a7 := newSetFromStringsArray(
		t,
		[]string{"l"},
		func(a, b interface{}) bool {return a == b},
	)
	p := newSetFromKSetsArray(
		t,
		[]KSet{
			KSet{key: "A1", set: a1},
			KSet{key: "A2", set: a2},
			KSet{key: "A3", set: a3},
			KSet{key: "A4", set: a4},
			KSet{key: "A5", set: a5},
			KSet{key: "A6", set: a6},
			KSet{key: "A7", set: a7},
		},
		func(a, b interface{}) bool {return a.(KSet).key == b.(KSet).key},
	)
	c := newSetFromKSetsArray(
		t,
		[]KSet{
			KSet{key: "A6", set: a6},
			KSet{key: "A2", set: a2},
			KSet{key: "A1", set: a1},
			KSet{key: "A3", set: a3},
		},
		func(a, b interface{}) bool {return a.(KSet).key == b.(KSet).key},
	)

	covering, err := cover(s, p)
	switch {
	case err != nil:
		t.Fatal("Error is not nil: ", err)
	case !covering.IsEqual(c):
		t.Error("Covering is not c: ", covering)
	}
}
