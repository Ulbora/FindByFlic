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
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	usession "github.com/Ulbora/go-better-sessions"
)

func TestHandler_HandleDcartIndex(t *testing.T) {
	var h Handler
	h.Templates = template.Must(template.ParseFiles("dcartIndex.html"))
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
	h.HandleDcartIndex(w, r)
}

func TestHandler_HandleDcartCb(t *testing.T) {
	var h Handler
	dcu := new(dcd.DCartUser)
	dcu.Action = "Activate"
	dcu.PublicKey = "123456"
	dcu.SecureURL = "http://someurl"
	dcu.TokenKey = "123456"
	aJSON, _ := json.Marshal(dcu)

	r, _ := http.NewRequest("POST", "/challenge", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleDcartCb(w, r)

}
