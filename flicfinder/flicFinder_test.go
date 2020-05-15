package flicfinder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

func TestFlicFinder_FindFlicListByZip(t *testing.T) {
	var ff FlicFinder
	var p px.MockGoProxy
	var res1 http.Response
	//res.StatusCode = 200
	res1.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"id":"123", "licenseName":"Some Lic Name", "businessName":"Some Bus", "premiseAddress":"Some Address"}]`))
	p.MockResp = &res1
	p.MockDoSuccess1 = true
	p.MockRespCode = 200
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

func TestFlicFinder_FindFlicListByZipFail(t *testing.T) {
	var ff FlicFinder
	var p px.MockGoProxy
	var res1 http.Response
	//res.StatusCode = 200
	res1.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"id":"123", "licenseName":"Some Lic Name", "businessName":"Some Bus", "premiseAddress":"Some Address"}]`))
	p.MockResp = &res1
	p.MockDoSuccess1 = false
	p.MockRespCode = 400
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

func TestFlicFinder_FindFlicByID(t *testing.T) {
	var ff FlicFinder
	var p px.MockGoProxy
	var res1 http.Response
	//res.StatusCode = 200
	res1.Body = ioutil.NopCloser(bytes.NewBufferString(`{"id":"123", "licenseName":"Some Lic Name", "businessName":"Some Bus", "premiseAddress":"Some Address"}`))
	p.MockResp = &res1
	p.MockDoSuccess1 = true
	p.MockRespCode = 200
	ff.Proxy = &p

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ff.Log = &l
	ff.FlicURL = "http://localhost:3000"
	ff.APIKey = "61616dfggdf5g64gf4"

	f := ff.GetNew()
	res := f.FindFlicByID("123")
	fmt.Println("res in get by id: ", *res)
	if res.Key != "123" {
		t.Fail()
	}
}

func TestFlicFinder_FindFlicByIDFail(t *testing.T) {
	var ff FlicFinder
	var p px.MockGoProxy
	var res1 http.Response
	//res.StatusCode = 200
	res1.Body = ioutil.NopCloser(bytes.NewBufferString(`{"id":"123", "licenseName":"Some Lic Name", "businessName":"Some Bus", "premiseAddress":"Some Address"}`))
	p.MockResp = &res1
	p.MockDoSuccess1 = false
	p.MockRespCode = 400
	ff.Proxy = &p

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ff.Log = &l
	ff.FlicURL = "http://localhost:3000"
	ff.APIKey = "61616dfggdf5g64gf4"

	f := ff.GetNew()
	res := f.FindFlicByID("123")
	fmt.Println("res in get by id: ", *res)
	if res.Key != "123" {
		t.Fail()
	}
}
