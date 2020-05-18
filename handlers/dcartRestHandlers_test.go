package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"

	dcd "github.com/Ulbora/FindByFlic/dbdelegate"
	ffl "github.com/Ulbora/FindByFlic/fflfinder"
	flc "github.com/Ulbora/FindByFlic/flicfinder"
	lg "github.com/Ulbora/Level_Logger"
	"github.com/gorilla/mux"

	//"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/Ulbora/dcartapi"
)

var dcDelResth dcd.DCartDeligate
var finderh ffl.Finder
var hh FlicHandler

// func TestHandlerRest_initFFLUserList(t *testing.T) {
// 	var mdb mydb.MyDB
// 	mdb.Host = "localhost:3306"
// 	mdb.User = "admin"
// 	mdb.Password = "admin"
// 	mdb.Database = "dcart_flic"
// 	dcDelRest.DB = &mdb
// 	h.FindFFLDCart = &dcDelRest
// 	suc := dcDelRest.DB.Connect()
// 	if !suc {
// 		t.Fail()
// 	}
// }

// func TestHandlerRest_initFFLList(t *testing.T) {
// 	var mdb mydb.MyDB
// 	mdb.Host = "localhost:3306"
// 	mdb.User = "admin"
// 	mdb.Password = "admin"
// 	mdb.Database = "ffl_list_10012018"
// 	finder.DB = &mdb
// 	h.FFLFinder = &finder
// 	suc := finder.DB.Connect()
// 	if !suc {
// 		t.Fail()
// 	}
// }

func TestHandlerRest_HandleFFLListFail(t *testing.T) {
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
	h := fh.GetNew()
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"zip": "30132",
	}
	r = mux.SetURLVars(r, vars)
	r.Header.Set("SecureURL", "http://someurl22")
	w := httptest.NewRecorder()
	h.HandleFFLList(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLList(t *testing.T) {
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
	h := fh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"zip": "30132",
	}
	r = mux.SetURLVars(r, vars)
	r.Header.Set("SecureURL", "http://someurl2")
	w := httptest.NewRecorder()
	h.HandleFFLList(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLListBadReq(t *testing.T) {
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
	h := fh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"zip": "30132",
	}
	r = mux.SetURLVars(r, vars)
	r.Header.Set("SecureURL", "http://someurl2")
	w := httptest.NewRecorder()
	h.HandleFFLList(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLGet(t *testing.T) {
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
	h := fh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "301",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("SecureURL", "http://someurl2")

	w := httptest.NewRecorder()
	h.HandleFFLGet(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLGetAuth(t *testing.T) {
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
	h := fh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "301",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("SecureURL", "http://someurl2")

	w := httptest.NewRecorder()
	h.HandleFFLGet(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLGetBadReq(t *testing.T) {
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
	h := fh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"id": "301",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("SecureURL", "http://someurl2")

	w := httptest.NewRecorder()
	h.HandleFFLGet(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddress(t *testing.T) {
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressMedia(t *testing.T) {
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressAuth(t *testing.T) {
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressBadStatus(t *testing.T) {
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

	var dapi api.MockAPI
	var odr api.Order
	odr.OrderID = 5
	odr.InvoiceNumber = 123
	odr.InvoiceNumberPrefix = "1001"
	dapi.MockOrder = &odr
	var shr api.ShipmentResponse
	shr.Status = "202"
	var shrs []api.ShipmentResponse
	shrs = append(shrs, shr)
	dapi.MockShipmentRes = &shrs
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressLicName(t *testing.T) {
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
	f1.LicName = "test bus"
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressLicNameOrderId(t *testing.T) {
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
	f1.LicName = "test bus"
	f1.Key = "123"
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressBadReq1(t *testing.T) {
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"", "invoice":"1000456"}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressBadReq(t *testing.T) {
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
	var secureURL = "http://someUrl"

	dapi.PrivateKey = "testKey"

	fh.DcartAPI = &dapi

	h := fh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"fflId":"1234", "invoice":"1000456"`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

// func TestHandlerRest_HandleFFLAddAddressMedia(t *testing.T) {
// 	var dapi api.API
// 	var secureURL string
// 	if len(os.Args) > 3 {
// 		privateKey := os.Args[2]
// 		secureURL = os.Args[3]
// 		//token := os.Args[3]
// 		//secureURL = os.Args[3]
// 		dapi.PrivateKey = privateKey
// 		//dapi.Token = token
// 		//dapi.SecureURL = secureURL

// 		log.Println("privateKey: ", privateKey)
// 		//log.Println("token: ", token)
// 		//log.Println("secureURL: ", secureURL)
// 		h.DcartAPI = &dapi
// 	}
// 	var req AddressRequest
// 	req.FFLID = "302"
// 	req.Invoice = "1054"
// 	aJSON, _ := json.Marshal(req)
// 	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
// 	//r.Header.Set("Content-Type", "application/json")
// 	r.Header.Set("SecureURL", secureURL)
// 	w := httptest.NewRecorder()
// 	h.HandleFFLAddAddress(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 415 {
// 		t.Fail()
// 	}
// }

// func TestHandlerRest_HandleFFLAddAddressReq(t *testing.T) {
// 	var dapi api.API
// 	var secureURL string
// 	if len(os.Args) > 3 {
// 		privateKey := os.Args[2]
// 		secureURL = os.Args[3]
// 		//token := os.Args[3]
// 		//secureURL = os.Args[3]
// 		dapi.PrivateKey = privateKey
// 		//dapi.Token = token
// 		//dapi.SecureURL = secureURL

// 		log.Println("privateKey: ", privateKey)
// 		//log.Println("token: ", token)
// 		//log.Println("secureURL: ", secureURL)
// 		h.DcartAPI = &dapi
// 	}
// 	var req AddressRequest
// 	//req.FFLID = "302"
// 	req.Invoice = "1054"
// 	aJSON, _ := json.Marshal(req)
// 	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	r.Header.Set("SecureURL", secureURL)
// 	w := httptest.NewRecorder()
// 	h.HandleFFLAddAddress(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 400 {
// 		t.Fail()
// 	}
// }

// func TestHandlerRest_HandleFFLAddAddressReq2(t *testing.T) {
// 	var dapi api.API
// 	var secureURL string
// 	if len(os.Args) > 3 {
// 		privateKey := os.Args[2]
// 		secureURL = os.Args[3]
// 		//token := os.Args[3]
// 		//secureURL = os.Args[3]
// 		dapi.PrivateKey = privateKey
// 		//dapi.Token = token
// 		//dapi.SecureURL = secureURL

// 		log.Println("privateKey: ", privateKey)
// 		//log.Println("token: ", token)
// 		//log.Println("secureURL: ", secureURL)
// 		h.DcartAPI = &dapi
// 	}
// 	var req AddressRequest
// 	req.FFLID = "302a"
// 	req.Invoice = "1054"
// 	aJSON, _ := json.Marshal(req)
// 	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	r.Header.Set("SecureURL", secureURL)
// 	w := httptest.NewRecorder()
// 	h.HandleFFLAddAddress(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 200 {
// 		t.Fail()
// 	}
// }

// func TestHandlerRest_HandleFFLAddAddressAuth(t *testing.T) {
// 	var dapi api.API
// 	//var secureURL string
// 	if len(os.Args) > 3 {
// 		privateKey := os.Args[2]
// 		//secureURL = os.Args[3]
// 		//token := os.Args[3]
// 		//secureURL = os.Args[3]
// 		dapi.PrivateKey = privateKey
// 		//dapi.Token = token
// 		//dapi.SecureURL = secureURL

// 		log.Println("privateKey: ", privateKey)
// 		//log.Println("token: ", token)
// 		//log.Println("secureURL: ", secureURL)
// 		h.DcartAPI = &dapi
// 	}
// 	var req AddressRequest
// 	req.FFLID = "302"
// 	req.Invoice = "1054"
// 	aJSON, _ := json.Marshal(req)
// 	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	r.Header.Set("SecureURL", "some.net")
// 	w := httptest.NewRecorder()
// 	h.HandleFFLAddAddress(w, r)
// 	fmt.Println("body: ", w.Code)
// 	if w.Code != 200 {
// 		t.Fail()
// 	}
// }

// func TestHandlerRest_close(t *testing.T) {
// 	suc := dcDelRest.DB.Close()
// 	fmt.Println("closing db")
// 	if !suc {
// 		t.Fail()
// 	}
// }

// func TestHandlerRest_close2(t *testing.T) {
// 	suc := finder.DB.Close()
// 	fmt.Println("closing db")
// 	if !suc {
// 		t.Fail()
// 	}
// }
