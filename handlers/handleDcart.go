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
	"encoding/json"
	"io/ioutil"
	//dbi "github.com/Ulbora/dbinterface"
	//ffl "FindByFlic/fflfinder"
	"log"
	"net/http"
)

//HandleDcartIndex HandleDcartIndex
func (h *Handler) HandleDcartIndex(w http.ResponseWriter, r *http.Request) {
	// h.Sess.InitSessionStore(w, r)
	// var order = r.URL.Query().Get("order")
	// var carturl = r.URL.Query().Get("carturl")
	// session := h.getSession(w, r)
	// session.Values["order"] = order
	// session.Values["carturl"] = carturl
	// serr := session.Save(r, w)
	// if serr != nil {
	// 	log.Println("Session Err:", serr)
	// }

	// var p PageParams
	// p.Order = order
	// p.CartURL = carturl
	// // p.URL = secureURL
	// // p.Key = privateKey
	// // p.Token = token
	// h.Templates.ExecuteTemplate(w, "dcartIndex.html", &p)
	h.Templates.ExecuteTemplate(w, "dcartIndex.html", nil)
}

//HandleDcartConfig HandleDcartConfig
func (h *Handler) HandleDcartConfig(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	h.Templates.ExecuteTemplate(w, "dcartConfig.html", nil)
}

//HandleDcartCb HandleDcartCb
func (h *Handler) HandleDcartCb(w http.ResponseWriter, r *http.Request) {
	var rtn bool
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("body err: ", err)
	} else {
		log.Println("json from dcart at start:", string(b))
	}

	log.Println("service call from 3dcart on callback-------")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		//var dcReg dcd.DCartUser
		//log.Println("body from dcart :", r.Body)
		dcReg := new(dcd.DCartUser)
		json.Unmarshal(b, dcReg)
		// decoder := json.NewDecoder(r.Body)
		// err := decoder.Decode(&dcReg)
		// if err != nil {
		// 	log.Println("Decode error: ", err.Error())
		// 	//http.Error(w, err.Error(), http.StatusBadRequest)
		// }

		if dcReg.Action == "AUTHORIZE" {
			rtn, _ = h.FindFFLDCart.AddUser(dcReg)
		} else if dcReg.Action == "REMOVE" {
			rtn = h.FindFFLDCart.RemoveUser(dcReg)
		}

		jsn, err := json.Marshal(dcReg)
		if err != nil {
			log.Println("marshal err: ", err)
		}
		log.Println("json from dcart at end:", string(jsn))
		if rtn {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

//HandleDcartFindFFL HandleDcartFindFFL
func (h *Handler) HandleDcartFindFFL(w http.ResponseWriter, r *http.Request) {
	zip := r.FormValue("zip")
	res := h.FFLFinder.FindFFL(zip)
	var pg FFLPageParams
	pg.FFLList = res
	pg.Zip = zip
	//log.Println("ffl list: ", pg.FFLList)
	h.Templates.ExecuteTemplate(w, "dcartAddFfl.html", &pg)
}

//HandleDcartChooseFFL HandleDcartChooseFFL
func (h *Handler) HandleDcartChooseFFL(w http.ResponseWriter, r *http.Request) {
	licNum := r.FormValue("id")
	zip := r.FormValue("zip")
	res := h.FFLFinder.GetFFL(licNum)
	var pg FFLPageParams
	pg.FFL = res
	pg.Zip = zip
	//log.Println("ffl: ", pg.FFL)
	h.Templates.ExecuteTemplate(w, "dcartChosenFfl.html", &pg)
}

//HandleDcartShipFFL HandleDcartShipFFL
func (h *Handler) HandleDcartShipFFL(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	licNum := r.FormValue("id")
	//log.Println("licNum in ship: ", licNum)
	//name := r.FormValue("name")
	//log.Println("name in ship: ", name)
	//address := r.FormValue("address")
	// log.Println("address in ship: ", address)
	// log.Println("order: ", session.Values["order"])
	// log.Println("carturl: ", session.Values["carturl"])
	res := h.FFLFinder.GetFFL(licNum)
	//var theFFL interface{}
	//theFFL = *res
	session.Values["fflLic"] = licNum
	//session.Values["fflName"] = name
	//session.Values["fflAddress"] = address
	serr := session.Save(r, w)
	if serr != nil {
		log.Println("Session Err:", serr)
	}
	var pg FFLPageParams
	pg.FFL = res
	//pg.Name = name
	//pg.Address = address
	//log.Println("shipTo: ", session.Values["fflLic"])
	//log.Println("ffl: ", pg.FFL)
	h.Templates.ExecuteTemplate(w, "dcartShippedFfl.html", &pg)
}

//HandleDcartShipFFLAddress HandleDcartShipFFLAddress
func (h *Handler) HandleDcartShipFFLAddress(w http.ResponseWriter, r *http.Request) {
	//h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	//licNum := r.FormValue("id")
	//log.Println("licNum in ship: ", licNum)
	//name := r.FormValue("name")
	//log.Println("name in ship: ", name)
	//address := r.FormValue("address")
	//log.Println("address in ship: ", address)
	var order = r.URL.Query().Get("order")
	var carturl = r.URL.Query().Get("carturl")
	log.Println("order in ffl address: ", order)
	log.Println("carturl in ffl address: ", carturl)
	//log.Println("shipTo: ", session.Values["shipTo"])
	var pg FFLPageParams
	fflLic := session.Values["fflLic"]
	var licNum string
	if fflLic != nil {
		licNum = fflLic.(string)
		log.Println("ffl lic in address: ", licNum)
		res := h.FFLFinder.GetFFL(licNum)
		pg.FFL = res
		//log.Println("shipTo in ffl in address: ", res)
	}

	//session.Values["shipTo"] = res
	//session.Save(r, w)

	//pg.Name = name
	//pg.Address = address
	//log.Println("shipTo in ffl lic in address: ", session.Values["fflLic"])
	//log.Println("shipTo in ffl name: ", session.Values["fflName"])
	//log.Println("shipTo in ffl address: ", session.Values["fflAddress"])

	//log.Println("ffl: ", pg.FFL)
	h.Templates.ExecuteTemplate(w, "dcartShipFflAddress.html", &pg)
}
