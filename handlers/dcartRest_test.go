package handlers

import (
	dcd "FindByFlic/dbdelegate"
	ffl "FindByFlic/fflfinder"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	//"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	mydb "github.com/Ulbora/dbinterface/mysql"
	api "github.com/Ulbora/dcartapi"
)

var dcDelRest dcd.DCartDeligate
var finder ffl.Finder
var h Handler

func TestHandlerRest_initFFLUserList(t *testing.T) {
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"
	dcDelRest.DB = &mdb
	h.FindFFLDCart = &dcDelRest
	suc := dcDelRest.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestHandlerRest_initFFLList(t *testing.T) {
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "ffl_list_10012018"
	finder.DB = &mdb
	h.FFLFinder = &finder
	suc := finder.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLListFail(t *testing.T) {
	r, _ := http.NewRequest("GET", "/ffllist?zip=30132", nil)
	r.Header.Set("SecureURL", "http://someurl22")
	w := httptest.NewRecorder()
	h.HandleFFLList(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/ffllist?zip=30132", nil)
	r.Header.Set("SecureURL", "http://someurl2")
	w := httptest.NewRecorder()
	h.HandleFFLList(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/ffllist?id=301", nil)
	r.Header.Set("SecureURL", "http://someurl2")
	w := httptest.NewRecorder()
	h.HandleFFLGet(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLGetBadReq(t *testing.T) {
	r, _ := http.NewRequest("GET", "/ffllist?id=301w", nil)
	r.Header.Set("SecureURL", "http://someurl2")
	w := httptest.NewRecorder()
	h.HandleFFLGet(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLGetBadUser(t *testing.T) {
	r, _ := http.NewRequest("GET", "/ffllist?id=301", nil)
	r.Header.Set("SecureURL", "http://someurl22")
	w := httptest.NewRecorder()
	h.HandleFFLGet(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddress(t *testing.T) {
	var dapi api.API
	var secureURL string
	if len(os.Args) > 3 {
		privateKey := os.Args[2]
		secureURL = os.Args[3]
		//token := os.Args[3]
		//secureURL = os.Args[3]
		dapi.PrivateKey = privateKey
		//dapi.Token = token
		//dapi.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		//log.Println("token: ", token)
		//log.Println("secureURL: ", secureURL)
		h.DcartAPI = &dapi
	}
	var req AddressRequest
	req.FFLID = "302"
	req.Invoice = "1054"
	aJSON, _ := json.Marshal(req)
	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
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
	var dapi api.API
	var secureURL string
	if len(os.Args) > 3 {
		privateKey := os.Args[2]
		secureURL = os.Args[3]
		//token := os.Args[3]
		//secureURL = os.Args[3]
		dapi.PrivateKey = privateKey
		//dapi.Token = token
		//dapi.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		//log.Println("token: ", token)
		//log.Println("secureURL: ", secureURL)
		h.DcartAPI = &dapi
	}
	var req AddressRequest
	req.FFLID = "302"
	req.Invoice = "1054"
	aJSON, _ := json.Marshal(req)
	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressReq(t *testing.T) {
	var dapi api.API
	var secureURL string
	if len(os.Args) > 3 {
		privateKey := os.Args[2]
		secureURL = os.Args[3]
		//token := os.Args[3]
		//secureURL = os.Args[3]
		dapi.PrivateKey = privateKey
		//dapi.Token = token
		//dapi.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		//log.Println("token: ", token)
		//log.Println("secureURL: ", secureURL)
		h.DcartAPI = &dapi
	}
	var req AddressRequest
	//req.FFLID = "302"
	req.Invoice = "1054"
	aJSON, _ := json.Marshal(req)
	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressReq2(t *testing.T) {
	var dapi api.API
	var secureURL string
	if len(os.Args) > 3 {
		privateKey := os.Args[2]
		secureURL = os.Args[3]
		//token := os.Args[3]
		//secureURL = os.Args[3]
		dapi.PrivateKey = privateKey
		//dapi.Token = token
		//dapi.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		//log.Println("token: ", token)
		//log.Println("secureURL: ", secureURL)
		h.DcartAPI = &dapi
	}
	var req AddressRequest
	req.FFLID = "302a"
	req.Invoice = "1054"
	aJSON, _ := json.Marshal(req)
	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", secureURL)
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_HandleFFLAddAddressAuth(t *testing.T) {
	var dapi api.API
	//var secureURL string
	if len(os.Args) > 3 {
		privateKey := os.Args[2]
		//secureURL = os.Args[3]
		//token := os.Args[3]
		//secureURL = os.Args[3]
		dapi.PrivateKey = privateKey
		//dapi.Token = token
		//dapi.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		//log.Println("token: ", token)
		//log.Println("secureURL: ", secureURL)
		h.DcartAPI = &dapi
	}
	var req AddressRequest
	req.FFLID = "302"
	req.Invoice = "1054"
	aJSON, _ := json.Marshal(req)
	r, _ := http.NewRequest("POST", "/ffllist", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("SecureURL", "some.net")
	w := httptest.NewRecorder()
	h.HandleFFLAddAddress(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandlerRest_close(t *testing.T) {
	suc := dcDelRest.DB.Close()
	fmt.Println("closing db")
	if !suc {
		t.Fail()
	}
}

func TestHandlerRest_close2(t *testing.T) {
	suc := finder.DB.Close()
	fmt.Println("closing db")
	if !suc {
		t.Fail()
	}
}
