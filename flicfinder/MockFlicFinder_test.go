package flicfinder

import (
	"testing"
)

func TestMockFlicFinder_FindFlicListByZip(t *testing.T) {
	var ff MockFlicFinder
	var flics []FlicList
	var f1 FlicList
	f1.BusName = "test bus"
	flics = append(flics, f1)
	ff.MockFlicList = &flics
	f := ff.GetNew()
	flist := f.FindFlicListByZip("12345")
	if len(*flist) == 0 {
		t.Fail()
	}
}

func TestMockFlicFinder_FindFlicByID(t *testing.T) {
	var ff MockFlicFinder
	var flic Flic
	flic.Key = "123455555"
	ff.MockFlic = &flic

	f := ff.GetNew()
	fl := f.FindFlicByID("123455555")
	if fl.Key != "123455555" {
		t.Fail()
	}
}
