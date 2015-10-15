package simple_list_elem

import "testing"

func TestSimpleIntListElem(t *testing.T) {
	var i = 10
	var list_elem = SimpleListElem{
		Data: i,
	}
	value, found := list_elem.Data.(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 10 {
		t.Error("data value incorrect")
	}
}

func TestSimpleStrListElem(t *testing.T) {
	var str = "hello"
	var list_elem = SimpleListElem{
		Data: str,
	}
	value, found := list_elem.Data.(string)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != "hello" {
		t.Error("data value incorrect")
	}
}
