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

func newFilledSet(t *testing.T, offset int, size int) Set {
	set := newEmptySet(t)
	limit := offset + size
	for i := offset; i < limit; i++ {
		set.Insert(i)
	}
	if set.Size() != size {
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
	set := newFilledSet(t, 0, 5)

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
	set1 := newFilledSet(t, 0, 5)
	set2 := newFilledSet(t, 3, 5)

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

func TestSetIntersection(t *testing.T) {
	set1 := newFilledSet(t, 0, 5)
	set2 := newFilledSet(t, 3, 5)

	intersection, err := set1.Intersection(set2)
	switch {
	case err != nil:
		t.Fatal("Error is not nil: ", err)
	case intersection.Size() != 2:
		t.Error("Size is not 2: ", intersection.Size())
	case !intersection.IsMember(3):
		t.Error("3 is not a member")
	case !intersection.IsMember(4):
		t.Error("4 is not a member")
	}
}

func TestSetDifference(t *testing.T) {
	set1 := newFilledSet(t, 0, 5)
	set2 := newFilledSet(t, 3, 5)

	diff, err := set1.Difference(set2)
	switch {
	case err != nil:
		t.Fatal("Error is not nil: ", err)
	case diff.Size() != 3:
		t.Error("Size is not 3: ", diff.Size())
	case diff.IsMember(3):
		t.Error("3 is a member")
	case diff.IsMember(4):
		t.Error("4 is a member")
	case !diff.IsMember(2):
		t.Error("2 is not a member")
	case !diff.IsMember(0):
		t.Error("0 is not a member")
	}
}

func TestIsMember(t *testing.T) {
	set := newFilledSet(t, 0, 5)

	switch {
	case set.IsMember(8):
		t.Error("IsMember returns true for 8")
	case !set.IsMember(4):
		t.Error("IsMember returns false for 4")
	}
}

func TestIsSubset(t *testing.T) {
	set1 := newFilledSet(t, 2, 4)
	set2 := newFilledSet(t, 0, 8)

	switch {
	case !set1.IsSubset(set2):
		t.Error("set1 is not a subset of set2")
	case set2.IsSubset(set1):
		t.Error("set2 is a subset of set1")
	}
}