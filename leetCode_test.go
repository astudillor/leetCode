package leetCode

import "testing"

func TestEmptyListPrint(t *testing.T) {
	// t.Fatal("not implemented")
	want := "-> //"
	var nodeNil *ListNode
	if got := nodeNil.String(); got != want {
		t.Errorf("nodeNil.String() = %q, want %q", got, want)
	}
}
