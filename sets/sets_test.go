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

func newFilledSet(t *testing.T) Set {
	set := newEmptySet(t)
	for i := 0; i < 5; i++ {
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
	set := newFilledSet(t)

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

func TestIsMember(t *testing.T) {
	set := newFilledSet(t)

	switch {
	case set.IsMember(8):
		t.Error("IsMember returns true for 8")
	case !set.IsMember(4):
		t.Error("IsMember returns false for 4")
	}
}