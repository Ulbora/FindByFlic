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
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	del "github.com/Ulbora/FindByFlic/dbdelegate"
	flc "github.com/Ulbora/FindByFlic/flicfinder"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/dcartapi"
	usession "github.com/Ulbora/go-better-sessions"
	"github.com/gorilla/sessions"
)

//Handler Handler
type Handler interface {
	HandleIndex(w http.ResponseWriter, r *http.Request)

	HandleDcartIndex(w http.ResponseWriter, r *http.Request)
	HandleDcartConfig(w http.ResponseWriter, r *http.Request)
	HandleDcartCb(w http.ResponseWriter, r *http.Request)

	HandleDcartFindFFL(w http.ResponseWriter, r *http.Request)
	HandleDcartChooseFFL(w http.ResponseWriter, r *http.Request)
	HandleDcartShipFFL(w http.ResponseWriter, r *http.Request)
	HandleDcartShipFFLAddress(w http.ResponseWriter, r *http.Request)

	// //REST handers
	HandleFFLList(w http.ResponseWriter, r *http.Request)
	HandleFFLGet(w http.ResponseWriter, r *http.Request)
	HandleFFLAddAddress(w http.ResponseWriter, r *http.Request)

	SetLogLevel(w http.ResponseWriter, r *http.Request)
}

//FlicHandler FlicHandler
type FlicHandler struct {
	Templates    *template.Template
	DCartUserDel del.DCartUserDelegate
	Sess         usession.Session
	//FFLFinder    ffl.FFLFinder
	FlicFinder flc.Finder
	DcartAPI   api.DcartAPI
	Log        *lg.Logger
}

//PageParams PageParams
type PageParams struct {
	Order   string
	CartURL string
	URL     string
	Key     string
	Token   string
	Enabled string
}

//FFLPageParams FFLPageParams
type FFLPageParams struct {
	//FFLList *[]ffl.FFL
	//FFL     *ffl.FFL
	FFLList *[]flc.FlicList
	FFL     *flc.Flic
	Zip     string
	Name    string
	Address string
	NoFFL   bool
	Enabled string
}

//GetNew GetNew
func (h *FlicHandler) GetNew() Handler {
	return h
}

func (h *FlicHandler) getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := h.Sess.GetSession(r)
	if err != nil {
		h.Log.Debug("get session error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return session
}

//HandleIndex HandleIndex
func (h *FlicHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	var order = r.URL.Query().Get("order")
	var carturl = r.URL.Query().Get("carturl")
	session := h.getSession(w, r)
	session.Values["order"] = order
	session.Values["carturl"] = carturl
	serr := session.Save(r, w)
	h.Log.Debug("Save Session err: ", serr)
	// if serr != nil {
	// 	h.Log.Debug(serr)
	// }

	// 	SecureUrl: 3dcart merchant's Secure URL.
	// PrivateKey: Your application's private key.
	// Token: The 3dcart merchant's token.
	//<iframe src="https://localhost:8070?cart=3dcar&carturl=[store_url]"></iframe>
	//<iframe src="https://localhost:8070?cart=3dcar&carturl=[store_url]"></iframe>

	// secureURL := r.Header.Get("SecureUrl")
	// privateKey := r.Header.Get("PrivateKey")
	// token := r.Header.Get("Token")
	var p PageParams
	p.Order = order
	p.CartURL = carturl
	// p.URL = secureURL
	// p.Key = privateKey
	// p.Token = token
	h.Templates.ExecuteTemplate(w, "index.html", &p)
}

//CheckContent CheckContent
func (h *FlicHandler) checkContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		rtn = true
	}
	return rtn
}

//SetContentType SetContentType
func (h *FlicHandler) setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//ProcessBody ProcessBody
func (h *FlicHandler) processBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(obj)
		if err != nil {
			h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}
