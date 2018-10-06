package fflfinder

import (
	"testing"
)

var fnd FFLFinder

func TestMockFinder_FindFFL(t *testing.T) {
	var f MockFinder
	fnd = &f
	lst := fnd.FindFFL("12345")
	if len(*lst) != 6 {
		t.Fail()
	}
}

func TestMockFinder_GetFFL(t *testing.T) {
	var f MockFinder
	fnd = &f
	fl := fnd.GetFFL(1)
	if fl == nil {
		t.Fail()
	}
}
