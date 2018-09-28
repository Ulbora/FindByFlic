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
