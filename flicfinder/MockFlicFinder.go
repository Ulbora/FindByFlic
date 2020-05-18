package flicfinder

import (
	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

//MockFlicFinder MockFlicFinder
type MockFlicFinder struct {
	Proxy        px.Proxy
	APIKey       string
	FlicURL      string
	Log          *lg.Logger
	MockFlicList *[]FlicList
	MockFlic     *Flic
}

//GetNew GetNew
func (f *MockFlicFinder) GetNew() Finder {
	return f
}

//FindFlicListByZip FindFlicListByZip
func (f *MockFlicFinder) FindFlicListByZip(zip string) *[]FlicList {
	return f.MockFlicList
}

//FindFlicByID FindFlicByID
func (f *MockFlicFinder) FindFlicByID(id string) *Flic {
	return f.MockFlic
}
