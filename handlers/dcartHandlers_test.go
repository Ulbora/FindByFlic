/*
 Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2018 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	dcd "github.com/Ulbora/FindByFlic/dbdelegate"
	flc "github.com/Ulbora/FindByFlic/flicfinder"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/dcartapi"

	usession "github.com/Ulbora/go-better-sessions"
)

// var dcart dcd.FindFFLDCart
// var dcDel dcd.DCartDeligate
// var db dbi.Database

// func TestHandler_init(t *testing.T) {
// 	var mdb mydb.MyDB
// 	mdb.Host = "localhost:3306"
// 	mdb.User = "admin"
// 	mdb.Password = "admin"
// 	mdb.Database = "dcart_flic"
// 	db = &mdb
// 	dcDel.DB = db
// 	dcart = &dcDel
// 	suc := dcDel.DB.Connect()
// 	if !suc {
// 		t.Fail()
// 	}
// }

func TestHandler_HandleDcartIndex(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l
	ff.Log = &l

	//var dapi api.MockAPI
	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartIndex.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	fh.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)
	session.Save(r, w)
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h := fh.GetNew()
	h.HandleDcartIndex(w, r)

	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartIndexAuth(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = false
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartIndex.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	fh.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)
	session.Save(r, w)
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h := fh.GetNew()
	h.HandleDcartIndex(w, r)

	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartConfig(t *testing.T) {
	var fh FlicHandler
	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartConfig.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	fh.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h := fh.GetNew()
	h.HandleDcartConfig(w, r)

	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCb(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu
	dcDelRest.MockAddSuccess = true

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"Action":"AUTHORIZE", "PublicKey":"1000456", "TimeStamp": "12-25-2018 01:01:00", "TokenKey":"123456", "SecureURL":"http://someurl"}`))
	r, _ := http.NewRequest("POST", "/challenge", aJSON)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbRemove(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	// var dcu dcd.DCartUser
	// dcu.Action = "AUTHORIZE"
	// dcu.PublicKey = "12345"
	// dcu.SecureURL = "https://teststore.cdcart.com"
	// dcu.TokenKey = "555ggg11"
	// dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	// dcu.Enabled = true
	// dcDelRest.MockDcartUser = &dcu
	dcDelRest.MockRemoveSuccess = true

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"Action":"REMOVE", "PublicKey":"1000456", "TimeStamp": "12-25-2018 01:01:00", "TokenKey":"123456", "SecureURL":"http://someurl"}`))
	r, _ := http.NewRequest("POST", "/challenge", aJSON)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbBacReqAction(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	// var dcu dcd.DCartUser
	// dcu.Action = "AUTHORIZE"
	// dcu.PublicKey = "12345"
	// dcu.SecureURL = "https://teststore.cdcart.com"
	// dcu.TokenKey = "555ggg11"
	// dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	// dcu.Enabled = true
	// dcDelRest.MockDcartUser = &dcu
	dcDelRest.MockAddSuccess = true

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"Action":"", "PublicKey":"1000456", "TimeStamp": "12-25-2018 01:01:00", "TokenKey":"123456", "SecureURL":"http://someurl"}`))
	r, _ := http.NewRequest("POST", "/challenge", aJSON)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbMedia(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu
	dcDelRest.MockAddSuccess = true

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"Action":"AUTHORIZE", "PublicKey":"1000456", "TimeStamp": "12-25-2018 01:01:00", "TokenKey":"123456", "SecureURL":"http://someurl"}`))
	r, _ := http.NewRequest("POST", "/challenge", aJSON)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbBadJsonReq(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu
	dcDelRest.MockAddSuccess = true

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var f1 flc.Flic
	f1.BusName = "test bus"
	f1.Key = "123"
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"Action":"AUTHORIZE", "PublicKey":"1000456", "TimeStamp": "12-25-2018 01:01:00", "TokenKey":"123456", "SecureURL":"http://someurl"`))
	r, _ := http.NewRequest("POST", "/challenge", aJSON)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

// func TestHandler_HandleDcartIndexUrl(t *testing.T) {
// 	var h FlicHandler
// 	h.FindFFLDCart = dcart
// 	h.Templates = template.Must(template.ParseFiles("testhtml/dcartIndex.html"))
// 	//h.TokenMap = make(map[string]*oauth2.Token)
// 	var s usession.Session
// 	h.Sess = s
// 	r, _ := http.NewRequest("GET", "/challenge?carturl=http://someurl", nil)
// 	w := httptest.NewRecorder()
// 	h.Sess.InitSessionStore(w, r)
// 	session, _ := h.Sess.GetSession(r)
// 	session.Save(r, w)
// 	//var resp oauth2.Token
// 	//resp.AccessToken = "bbbnn"
// 	//h.TokenMap["123456"] = &resp
// 	h.HandleDcartIndex(w, r)

// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 200 {
// 		t.Fail()
// 	}
// }

// func TestHandler_HandleDcartCbMedia(t *testing.T) {
// 	var h FlicHandler
// 	h.FindFFLDCart = dcart
// 	dcu := new(dcd.DCartUser)
// 	dcu.Action = "AUTHORIZE"
// 	dcu.PublicKey = "123456"
// 	dcu.SecureURL = "http://someurl"
// 	dcu.TokenKey = "123456"
// 	dcu.TimeStamp = "12-25-2018 01:01:00"
// 	aJSON, _ := json.Marshal(dcu)

// 	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
// 	//r.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	h.HandleDcartCb(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 415 {
// 		t.Fail()
// 	}
// }

// func TestHandler_HandleDcartCbRemove(t *testing.T) {
// 	var h FlicHandler
// 	h.FindFFLDCart = dcart
// 	dcu := new(dcd.DCartUser)
// 	dcu.Action = "REMOVE"
// 	dcu.PublicKey = "123456"
// 	dcu.SecureURL = "http://someurl"
// 	dcu.TokenKey = "123456"
// 	dcu.TimeStamp = "12-25-2018 01:01:00"
// 	aJSON, _ := json.Marshal(dcu)

// 	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	h.HandleDcartCb(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 200 {
// 		t.Fail()
// 	}
// }

// func TestHandler_HandleDcartCbRemoveFail(t *testing.T) {
// 	var h FlicHandler
// 	h.FindFFLDCart = dcart
// 	dcu := new(dcd.DCartUser)
// 	dcu.Action = "REMOVE1"
// 	dcu.PublicKey = "123456"
// 	dcu.SecureURL = "http://someurl"
// 	dcu.TokenKey = "123456"
// 	dcu.TimeStamp = "12-25-2018 01:01:00"
// 	aJSON, _ := json.Marshal(dcu)

// 	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	h.HandleDcartCb(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 400 {
// 		t.Fail()
// 	}
// }

func TestHandler_HandleDcartFindFFL(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate
	// var dcu dcd.DCartUser
	// dcu.Action = "AUTHORIZE"
	// dcu.PublicKey = "12345"
	// dcu.SecureURL = "https://teststore.cdcart.com"
	// dcu.TokenKey = "555ggg11"
	// dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	// dcu.Enabled = true
	// dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var flics []flc.FlicList
	var f1 flc.FlicList
	f1.BusName = "test bus"
	flics = append(flics, f1)
	ff.MockFlicList = &flics
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartAddFfl.html"))
	//fh.FFLFinder = new(ffl.MockFinder)

	r, _ := http.NewRequest("GET", "/challenge?zip=12345", nil)
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartFindFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartFindFFLUse(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	var flics []flc.FlicList
	var f1 flc.FlicList
	f1.BusName = "test bus"
	flics = append(flics, f1)
	ff.MockFlicList = &flics
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartAddFfl.html"))
	//fh.FFLFinder = new(ffl.MockFinder)

	r, _ := http.NewRequest("POST", "/challenge?zip=12345&use=true", nil)
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartFindFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartChooseFFL(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	//var flics []flc.FlicList
	var f1 flc.Flic
	f1.BusName = "test bus"
	//flics = append(flics, f1)
	//ff.MockFlicList = &flics
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartChosenFfl.html"))
	//h.FFLFinder = new(ffl.MockFinder)

	r, _ := http.NewRequest("POST", "/challenge?zip=12345&id=1234&use=true", nil)
	w := httptest.NewRecorder()
	h := fh.GetNew()
	h.HandleDcartChooseFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

// // type TestShip struct{
// // 	ID string
// // 	Name string
// // 	Address string
// // }
func TestHandler_HandleDcartShipFFL(t *testing.T) {

	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder
	//var flics []flc.FlicList
	var f1 flc.Flic
	f1.BusName = "test bus"
	//flics = append(flics, f1)
	//ff.MockFlicList = &flics
	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi
	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))
	//h.FFLFinder = new(ffl.MockFinder)
	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("POST", "/challenge?id=12345&use=true", nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)
	session.Values["order"] = "AB1026"
	session.Values["carturl"] = "https://testcart.3dcart.com"
	session.Save(r, w)
	h := fh.GetNew()
	h.HandleDcartShipFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartShipFFLAddress(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder

	var f1 flc.Flic
	f1.BusName = "test bus"

	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))

	var secureURL string

	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)

	session.Values["fflLic"] = "12345"
	session.Save(r, w)

	h := fh.GetNew()
	h.HandleDcartShipFFLAddress(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartShipFFLAddressNoLicID(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder

	var f1 flc.Flic
	f1.BusName = "test bus"

	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))

	var secureURL string

	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	// session, _ := fh.Sess.GetSession(r)

	// session.Values["fflLic"] = "12345"
	// session.Save(r, w)

	h := fh.GetNew()
	h.HandleDcartShipFFLAddress(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartShipFFLAddressAuth(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = false
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder

	var f1 flc.Flic
	f1.BusName = "test bus"

	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))

	var secureURL string

	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)

	session.Values["fflLic"] = "12345"
	session.Save(r, w)

	h := fh.GetNew()
	h.HandleDcartShipFFLAddress(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartShipFFLAddressOrderErr(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder

	var f1 flc.Flic
	f1.BusName = "test bus"

	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	//odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))

	var secureURL string

	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)

	session.Values["fflLic"] = "12345"
	session.Save(r, w)

	h := fh.GetNew()
	h.HandleDcartShipFFLAddress(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartShipFFLAddressErr(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder

	var f1 flc.Flic
	f1.BusName = "test bus"

	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "204"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))

	var secureURL string

	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)

	session.Values["fflLic"] = "12345"
	session.Save(r, w)

	h := fh.GetNew()
	h.HandleDcartShipFFLAddress(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartShipFFLAddressNotBus(t *testing.T) {
	var fh FlicHandler
	var dcDelRest dcd.MockDCartDeligate

	var dcu dcd.DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	dcDelRest.MockDcartUser = &dcu

	fh.DCartUserDel = &dcDelRest

	var ff flc.MockFlicFinder

	var f1 flc.Flic
	//f1.BusName = "test bus"
	f1.LicName = "some person"

	ff.MockFlic = &f1
	fh.FlicFinder = &ff

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	dcDelRest.Log = &l

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "201"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	//	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	fh.Templates = template.Must(template.ParseFiles("testhtml/dcartShippedFfl.html"))

	var secureURL string

	var s usession.Session
	fh.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)

	session.Values["fflLic"] = "12345"
	session.Save(r, w)

	h := fh.GetNew()
	h.HandleDcartShipFFLAddress(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}
