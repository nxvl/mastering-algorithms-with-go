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

func (set Set) IsMember(data interface{}) bool {
	for member := set.Head(); member != nil; member = set.Next(member) {
		if set.match(data, set.Data(member)) {
			return true
		}
	}

	return false
}
