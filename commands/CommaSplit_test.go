package commands

import "testing"

func TestCommaSplit(t *testing.T) {
	str := "a, b, \t\t\tc,  d,e"
	str2 := "x"
	list := CommaSplit(str)
	list2 := CommaSplit(str2)

	if len(list) != 5 {
		t.Fail()
		return
	}
	if list[0] != "a" || list[1] != "b" || list[2] != "c" || list[3] != "d" || list[4] != "e" {
		t.Fail()
		return
	}
	if len(list2) != 1 {
		t.Fail()
		return
	}
	if len(CommaSplit("")) != 0 {
		t.Fail()
		return
	}
}
