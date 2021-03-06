package sets

import (
	list "github.com/nxvl/mastering-algorithms-with-go/linked_lists"
)

// Set struct. Inherits from SimpleLinkedList
type Set struct {
	list.SimpleLinkedList
	match func(interface{}, interface{}) bool
}

func NewSet(match func(interface{}, interface{}) bool) (Set, error) {
	return Set{match: match}, nil
}

func (set *Set) Insert(data interface{}) error {
	if set.IsMember(data) {
		return nil
	}
	return set.InsNext(set.Tail(), data)
}

func (set *Set) Remove(data interface{}) (interface{}, error) {
	var member, prev *list.SimpleListElem
	for member = set.Head(); member != nil; member = set.Next(member) {
		if set.match(data, set.Data(member)){
			break
		}
		prev = member
	}

	if member == nil {
		return nil, ErrNothingToRemove
	}

	return set.RemNext(prev)
}

func (set1 Set) Union(set2 Set) (Set, error) {
	setu, error := NewSet(set1.match)
	if error != nil {
		return *new(Set), error
	}

	for _, set := range([2]Set{set1, set2}) {
		for member := set.Head(); member != nil; member = set.Next(member) {
			data := set.Data(member)
			error = setu.Insert(data)
			if error != nil {
				return *new(Set), error
			}
		}
	}

	return setu, nil
}

func (set1 Set) Intersection(set2 Set) (Set, error) {
	seti, error := NewSet(set1.match)
	if error != nil {
		return *new(Set), error
	}

	for member:= set1.Head(); member != nil; member = set1.Next(member) {
		data := set1.Data(member)
		if set2.IsMember(data) {
			error = seti.Insert(data)
			if error != nil {
				return *new(Set), error
			}
		}
	}

	return seti, nil
}

func (set1 Set) Difference(set2 Set) (Set, error) {
	setd, error := NewSet(set1.match)
	if error != nil {
		return *new(Set), error
	}

	for member:= set1.Head(); member != nil; member = set1.Next(member) {
		data := set1.Data(member)
		if !set2.IsMember(data) {
			error = setd.Insert(data)
			if error != nil {
				return *new(Set), error
			}
		}
	}

	return setd, nil
}

func (set Set) IsMember(data interface{}) bool {
	for member := set.Head(); member != nil; member = set.Next(member) {
		if set.match(data, set.Data(member)) {
			return true
		}
	}

	return false
}

func (set1 Set) IsSubset(set2 Set) bool {
	if set1.Size() > set2.Size() {
		return false
	}

	for member := set1.Head(); member != nil; member = set1.Next(member) {
		data := set1.Data(member)
		if !set2.IsMember(data) {
			return false
		}
	}

	return true
}

func (set1 Set) IsEqual(set2 Set) bool {
	if set1.Size() == set2.Size() && set1.IsSubset(set2) {
		return true
	} else {
		return false
	}
}
