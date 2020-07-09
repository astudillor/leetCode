package leetCode

import "testing"

func TestEmptyListPrint(t *testing.T) {
	// t.Fatal("not implemented")
	want := "//"
	var nodeNil *ListNode
	if got := nodeNil.String(); got != want {
		t.Errorf("nodeNil.String() = %q, want %q", got, want)
	}
}

func TestInts2List(t *testing.T) {
	// t.Fatal("not implemented")
	want := "1 -> 2 -> 3 -> 6 -> //"
	head := Ints2List(1, 2, 3, 6)
	if got := head.String(); got != want {
		t.Errorf("nodeNil.String() = %q, want %q", got, want)
	}
}
