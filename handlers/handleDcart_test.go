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
	dcd "FindByFlic/dbdelegate"
	ffl "FindByFlic/fflfinder"
	"bytes"
	"encoding/json"
	"fmt"
	api "github.com/Ulbora/dcartapi"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	dbi "github.com/Ulbora/dbinterface"
	mydb "github.com/Ulbora/dbinterface/mysql"
	usession "github.com/Ulbora/go-better-sessions"
)

var dcart dcd.FindFFLDCart
var dcDel dcd.DCartDeligate
var db dbi.Database

func TestHandler_init(t *testing.T) {
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"
	db = &mdb
	dcDel.DB = db
	dcart = &dcDel
	suc := dcDel.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestHandler_HandleDcartIndex(t *testing.T) {
	var h Handler
	h.FindFFLDCart = dcart
	h.Templates = template.Must(template.ParseFiles("dcartIndex.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Save(r, w)
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h.HandleDcartIndex(w, r)

	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartConfig(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("dcartConfig.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h.HandleDcartConfig(w, r)

	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCb(t *testing.T) {
	var h Handler
	h.FindFFLDCart = dcart
	dcu := new(dcd.DCartUser)
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "123456"
	dcu.SecureURL = "http://someurl"
	dcu.TokenKey = "123456"
	dcu.TimeStamp = "12-25-2018 01:01:00"
	aJSON, _ := json.Marshal(dcu)

	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartIndexUrl(t *testing.T) {
	var h Handler
	h.FindFFLDCart = dcart
	h.Templates = template.Must(template.ParseFiles("dcartIndex.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?carturl=http://someurl", nil)
	w := httptest.NewRecorder()
	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Save(r, w)
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h.HandleDcartIndex(w, r)

	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbMedia(t *testing.T) {
	var h Handler
	h.FindFFLDCart = dcart
	dcu := new(dcd.DCartUser)
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "123456"
	dcu.SecureURL = "http://someurl"
	dcu.TokenKey = "123456"
	dcu.TimeStamp = "12-25-2018 01:01:00"
	aJSON, _ := json.Marshal(dcu)

	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbRemove(t *testing.T) {
	var h Handler
	h.FindFFLDCart = dcart
	dcu := new(dcd.DCartUser)
	dcu.Action = "REMOVE"
	dcu.PublicKey = "123456"
	dcu.SecureURL = "http://someurl"
	dcu.TokenKey = "123456"
	dcu.TimeStamp = "12-25-2018 01:01:00"
	aJSON, _ := json.Marshal(dcu)

	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartCbRemoveFail(t *testing.T) {
	var h Handler
	h.FindFFLDCart = dcart
	dcu := new(dcd.DCartUser)
	dcu.Action = "REMOVE1"
	dcu.PublicKey = "123456"
	dcu.SecureURL = "http://someurl"
	dcu.TokenKey = "123456"
	dcu.TimeStamp = "12-25-2018 01:01:00"
	aJSON, _ := json.Marshal(dcu)

	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleDcartCb(w, r)
	fmt.Println("body: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestHandler_HandleDcartFindFFL(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("dcartAddFfl.html"))
	h.FFLFinder = new(ffl.MockFinder)

	r, _ := http.NewRequest("GET", "/challenge?zip=12345", nil)
	w := httptest.NewRecorder()
	h.HandleDcartFindFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleDcartChooseFFL(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("dcartChosenFfl.html"))
	h.FFLFinder = new(ffl.MockFinder)

	r, _ := http.NewRequest("GET", "/challenge?zip=12345", nil)
	w := httptest.NewRecorder()
	h.HandleDcartChooseFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

// type TestShip struct{
// 	ID string
// 	Name string
// 	Address string
// }
func TestHandler_HandleDcartShipFFL(t *testing.T) {

	var h Handler
	h.Templates = template.Must(template.ParseFiles("dcartShippedFfl.html"))
	h.FFLFinder = new(ffl.MockFinder)
	var s usession.Session
	h.Sess = s

	r, _ := http.NewRequest("POST", "/challenge?id=12345&name=bobs guns&address=125 marietta, ga 12345", nil)
	w := httptest.NewRecorder()

	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	session.Values["order"] = "AB1026"
	session.Values["carturl"] = "https://testcart.3dcart.com"
	session.Save(r, w)
	h.HandleDcartShipFFL(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

//to test
//go test -coverprofile=coverage.out -args privateKey secureURL
func TestHandler_HandleDcartShipFFLAddress(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("dcartShippedFfl.html"))
	h.FFLFinder = new(ffl.MockFinder)
	h.FindFFLDCart = dcart
	var secureURL string
	var dapi api.API
	if len(os.Args) == 4 {
		privateKey := os.Args[2]
		//token := os.Args[3]
		secureURL = os.Args[3]
		dapi.PrivateKey = privateKey
		//dapi.Token = token
		//dapi.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		//log.Println("token: ", token)
		//log.Println("secureURL: ", secureURL)
		h.DcartAPI = &dapi
	}
	var s usession.Session
	h.Sess = s

	r, _ := http.NewRequest("GET", "/challenge?invoice=1041&carturl="+secureURL, nil)
	w := httptest.NewRecorder()

	h.Sess.InitSessionStore(w, r)
	session, _ := h.Sess.GetSession(r)
	var finder = new(ffl.MockFinder)
	f1 := finder.GetFFL(5)
	idstr := strconv.FormatInt(f1.ID, 10)
	session.Values["fflLic"] = idstr
	session.Save(r, w)
	//session.Values["carturl"] = "https://testcart.3dcart.com"
	//session.Save(r, w)
	h.HandleDcartShipFFLAddress(w, r)
	// body, _ := ioutil.ReadAll(w.Result().Body)
	// fmt.Println("ffl list Res body", string(body))
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_close(t *testing.T) {
	suc := dcDel.DB.Close()
	fmt.Println("closing db")
	if !suc {
		t.Fail()
	}
}
