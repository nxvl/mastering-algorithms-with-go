package sets

import (
	"testing"
)

func newEmptySet(t *testing.T) Set {
	set, err := NewSet(func(a, b interface{}) bool {return a == b})
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
	return set
}

func newFilledSet(t *testing.T, offset int) Set {
	set := newEmptySet(t)
	limit := offset + 5
	for i := offset; i < limit; i++ {
		set.Insert(i)
	}
	if set.Size() != 5 {
		t.Fatal("Wrong initial set size: ", set.Size())
	}
	return set
}

func TestSetInsert(t *testing.T) {
	set := newEmptySet(t)

	err := set.Insert(1)
	switch {
	case err != nil:
		t.Error("Error is not nil: ", err)
	case set.Size() != 1:
		t.Error("Wrong size: ", set.Size())
	}

	err = set.Insert(4)
	switch {
	case err != nil:
		t.Error("Error is not nil: ", err)
	case set.Size() != 2:
		t.Error("Wrong size: ", set.Size())
	}

	err = set.Insert(1)
	switch {
	case err != nil:
		t.Error("Error is not nil: ", err)
	case set.Size() != 2:
		t.Error("Wrong size: ", set.Size())
	}
}

func TestSetRemove(t *testing.T) {
	set := newFilledSet(t, 0)

	elem, err := set.Remove(2)
	switch {
	case err != nil:
		t.Fatal("Error is not nil: ", err)
	case set.Size() != 4:
		t.Error("Size is not 4: ", set.Size())
	case elem != 2:
		t.Error("Wrong elem value: ", elem)
	}

	elem, err = set.Remove(9)
	switch {
	case err != ErrNothingToRemove:
		t.Fatal("Error is not ErrNothingToRemove: ", err)
	case set.Size() != 4:
		t.Error("Size is not 4: ", set.Size())
	case elem != nil:
		t.Error("elem is not nil: ", elem)
	}
}

func TestSetUnion(t *testing.T) {
	set1 := newFilledSet(t, 0)
	set2 := newFilledSet(t, 3)

	union, err := set1.Union(set2)
	switch {
	case err != nil:
		t.Fatal("Error is not nil: ", err)
	case union.Size() != 8:
		t.Error("Size is not 8: ", union.Size())
	case !union.IsMember(7):
		t.Error("7 is not a member")
	case !union.IsMember(0):
		t.Error("0 is not a member")
	}
}

func TestIsMember(t *testing.T) {
	set := newFilledSet(t, 0)

	switch {
	case set.IsMember(8):
		t.Error("IsMember returns true for 8")
	case !set.IsMember(4):
		t.Error("IsMember returns false for 4")
	}
}