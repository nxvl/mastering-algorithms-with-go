package sets

import (
	"errors"
	"github.com/nxvl/mastering-algorithms-with-go/linked_lists"
)

type CoveringError error

var CanNotCover = errors.New("Can't cover set with given subsets")

type KSet struct {
	key interface{}
	set Set
}

func cover(members Set, subsets Set) (Set, error) {
	var intersection Set
	var subset, tmp_kset KSet
	var max_member, member *linked_list.SimpleListElem
	var max_size int

	// Init the covering
	covering, err := NewSet(subsets.match)

	if err != nil {
		return *new(Set), err
	}

	// Continue while there are non covered members and candidate subsets.
	for members.Size() > 0 && subsets.Size() > 0 {
		max_size = 0
		// Find the subset that covers the most members.
		for member = subsets.Head(); member != nil; member = subsets.Next(member) {
			tmp_kset = subsets.Data(member).(KSet)
			intersection, err = tmp_kset.set.Intersection(members)
			if err != nil {
				return *new(Set), err
			}

			if intersection.Size() > max_size {
				max_member = member
				max_size = intersection.Size()
			}
		}

		// A covering is not possible if there was no intersection.
		if max_size == 0 {
			return *new(Set), CanNotCover
		}

		// Insert the selected subset into the covering.
		subset = subsets.Data(max_member).(KSet)
		err = covering.Insert(subset)
		if err != nil {
			return *new(Set), err
		}

		// Remove each covered member from the set of noncovered members.
		tmp_kset = subsets.Data(max_member).(KSet)
		tmp_set := tmp_kset.set
		for member = tmp_set.Head(); member != nil; member = tmp_set.Next(member) {
			data := tmp_set.Data(member)
			_, err = members.Remove(data)
			if err != nil && err != ErrNothingToRemove{
				return *new(Set), err
			}
		}

		// Remove the subset for the set of candidate subsets.
		_, err = subsets.Remove(subset)
		if err != nil {
			return *new(Set), err
		}
		subset = *new(KSet)
	}

	// No covering is possible if there are still noncovered members.
	if members.Size() > 0 {
		return *new(Set), CanNotCover
	}

	return covering, nil
}
