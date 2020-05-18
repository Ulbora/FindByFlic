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
	"html/template"
	"net/http"
	"net/http/httptest"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	usession "github.com/Ulbora/go-better-sessions"
	//"github.com/gorilla/sessions"
)

func TestHandler_HandleIndex(t *testing.T) {
	var fh FlicHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l

	fh.Templates = template.Must(template.ParseFiles("testhtml/index.html"))
	//h.TokenMap = make(map[string]*oauth2.Token)
	var s usession.Session
	fh.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	fh.Sess.InitSessionStore(w, r)
	session, _ := fh.Sess.GetSession(r)
	session.Values["accessTokenKey"] = "123456"
	h := fh.GetNew()
	//var resp oauth2.Token
	//resp.AccessToken = "bbbnn"
	//h.TokenMap["123456"] = &resp
	h.HandleIndex(w, r)
}

func TestHandler_getSession(t *testing.T) {
	var fh FlicHandler
	var s usession.Session
	fh.Sess = s
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	fh.Sess.InitSessionStore(w, r)
	//h := fh.GetNew()
	ss := fh.getSession(w, r)
	if ss == nil {
		t.Fail()
	}
}

func TestHandler_getSessionFail(t *testing.T) {
	var fh FlicHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	fh.Log = &l
	//var s usession.Session
	//h.Sess = s
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	//h.Sess.InitSessionStore(w, r)
	ss := fh.getSession(w, r)
	if ss != nil {
		t.Fail()
	}
}
