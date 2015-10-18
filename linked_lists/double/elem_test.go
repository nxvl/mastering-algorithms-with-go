package double_list

import "testing"

func TestDoubleIntListElem(t *testing.T) {
	var i = 10
	var list_elem = DoubleListElem{
		data: i,
	}
	value, found := list_elem.data.(int)
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
		data: str,
	}
	value, found := list_elem.data.(string)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != "hello" {
		t.Error("data value incorrect")
	}
}
