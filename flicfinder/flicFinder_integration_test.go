// +build integration move to top

package flicfinder

import (
	"fmt"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

func TestFlicFinderi_FindFlicListByZip(t *testing.T) {
	var ff FlicFinder
	var p px.GoProxy
	ff.Proxy = &p

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ff.Log = &l
	ff.FlicURL = "http://localhost:3000"
	ff.APIKey = "61616dfggdf5g64gf4"

	f := ff.GetNew()
	res := f.FindFlicListByZip("30141")
	fmt.Println("res in get by zip: ", *res)
	if len(*res) == 0 {
		t.Fail()
	}

}

func TestFlicFinderi_FindFlicByID(t *testing.T) {
	var ff FlicFinder
	var p px.GoProxy
	ff.Proxy = &p

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ff.Log = &l
	ff.FlicURL = "http://localhost:3000"
	ff.APIKey = "61616dfggdf5g64gf4"

	f := ff.GetNew()
	res := f.FindFlicByID("158223021F05412")
	fmt.Println("res in get by id: ", *res)
	if res.Key != "158223021F05412" {
		t.Fail()
	}

}
