package double_list_elem

import "testing"

func TestDoubleIntListElem(t *testing.T) {
	var i = 10
	var list_elem = DoubleListElem{
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

func TestDoubleStrListElem(t *testing.T) {
	var str = "hello"
	var list_elem = DoubleListElem{
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
