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
	"io/ioutil"
	"strconv"

	dcd "github.com/Ulbora/FindByFlic/dbdelegate"

	//dbi "github.com/Ulbora/dbinterface"
	//ffl "FindByFlic/fflfinder"
	"log"
	"net/http"

	ffl "github.com/Ulbora/FindByFlic/fflfinder"
	api "github.com/Ulbora/dcartapi"
)

//HandleDcartIndex HandleDcartIndex
func (h *Handler) HandleDcartIndex(w http.ResponseWriter, r *http.Request) {
	//h.Sess.InitSessionStore(w, r)
	var carturl = r.URL.Query().Get("carturl")
	log.Println("carturl in ffl address: ", carturl)
	ures := h.FindFFLDCart.GetUser(carturl)
	log.Println("ures in index: ", ures)
	var pg PageParams
	if ures != nil && ures.Enabled {
		pg.Enabled = "true"
	} else {
		pg.Enabled = "false"
	}
	h.Templates.ExecuteTemplate(w, "dcartIndex.html", &pg)
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
	use := r.FormValue("use")
	var res = new([]ffl.FFL)
	if use == "true" {
		res = h.FFLFinder.FindFFL(zip)
	}
	var pg FFLPageParams
	pg.FFLList = res
	pg.Zip = zip
	pg.Enabled = use
	//log.Println("ffl list: ", pg.FFLList)
	h.Templates.ExecuteTemplate(w, "dcartAddFfl.html", &pg)
}

//HandleDcartChooseFFL HandleDcartChooseFFL
func (h *Handler) HandleDcartChooseFFL(w http.ResponseWriter, r *http.Request) {
	idstr := r.FormValue("id")
	zip := r.FormValue("zip")
	use := r.FormValue("use")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	res := h.FFLFinder.GetFFL(id)
	var pg FFLPageParams
	pg.FFL = res
	pg.Zip = zip
	pg.Enabled = use
	//log.Println("ffl: ", pg.FFL)
	h.Templates.ExecuteTemplate(w, "dcartChosenFfl.html", &pg)
}

//HandleDcartShipFFL HandleDcartShipFFL
func (h *Handler) HandleDcartShipFFL(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	idstr := r.FormValue("id")
	use := r.FormValue("use")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	//log.Println("licNum in ship: ", licNum)
	//name := r.FormValue("name")
	//log.Println("name in ship: ", name)
	//address := r.FormValue("address")
	// log.Println("address in ship: ", address)
	// log.Println("order: ", session.Values["order"])
	// log.Println("carturl: ", session.Values["carturl"])
	res := h.FFLFinder.GetFFL(id)
	//var theFFL interface{}
	//theFFL = *res
	session.Values["fflLic"] = idstr
	//session.Values["fflName"] = name
	//session.Values["fflAddress"] = address
	serr := session.Save(r, w)
	if serr != nil {
		log.Println("Session Err:", serr)
	}
	var pg FFLPageParams
	pg.FFL = res
	pg.Enabled = use
	//pg.Name = name
	//pg.Address = address
	//log.Println("shipTo: ", session.Values["fflLic"])
	//log.Println("ffl: ", pg.FFL)
	h.Templates.ExecuteTemplate(w, "dcartShippedFfl.html", &pg)
}

//HandleDcartShipFFLAddress HandleDcartShipFFLAddress
func (h *Handler) HandleDcartShipFFLAddress(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	//licNum := r.FormValue("id")
	//log.Println("licNum in ship: ", licNum)
	//name := r.FormValue("name")
	//log.Println("name in ship: ", name)
	//address := r.FormValue("address")
	//log.Println("address in ship: ", address)
	var invoice = r.URL.Query().Get("invoice")
	var carturl = r.URL.Query().Get("carturl")
	log.Println("invoice in ffl address: ", invoice)
	log.Println("carturl in ffl address: ", carturl)
	//log.Println("shipTo: ", session.Values["shipTo"])
	var pg FFLPageParams
	fflLicID := session.Values["fflLic"]
	log.Println("fflLic: ", fflLicID)
	//var licNum string
	if fflLicID != nil {
		//idstr := fflLicID.(string)
		idstr := fflLicID.(string)
		///id := int64(idint)
		log.Println("ffl id: ", idstr)
		id, _ := strconv.ParseInt(idstr, 10, 64)
		ures := h.FindFFLDCart.GetUser(carturl)
		log.Println("user found before if: ", ures)
		if ures.Enabled {
			log.Println("user found: ", ures)
			//licNum = fflLic.(string)
			log.Println("ffl lic in address: ", id)
			res := h.FFLFinder.GetFFL(id)
			pg.FFL = res

			odr := h.DcartAPI.GetOrder(invoice, ures.SecureURL, ures.TokenKey)
			if odr.OrderID != 0 {
				var s api.Shipment
				s.ShipmentID = 0
				s.ShipmentFirstName = "FFL"
				s.ShipmentLastName = "Lic # " + res.LicRegn + res.LicDist + res.LicCnty + res.LicType + res.LicXprdte + res.LicSeqn
				if res.BusinessName != "NULL" {
					s.ShipmentCompany = res.BusinessName
				} else {
					s.ShipmentCompany = res.LicenseName
				}

				s.ShipmentAddress = res.PremiseStreet
				s.ShipmentCity = res.PremiseCity
				s.ShipmentState = res.PremiseState
				s.ShipmentZipCode = res.PremiseZipCode
				s.ShipmentCountry = "USA"
				s.ShipmentPhone = res.VoicePhone
				s.ShipmentTax = 0
				s.ShipmentWeight = 1
				s.ShipmentTrackingCode = ""
				s.ShipmentUserID = ""
				s.ShipmentNumber = 1
				s.ShipmentAddressTypeID = 0
				var oid = strconv.FormatInt(odr.OrderID, 10)
				sres := h.DcartAPI.AddShippingAddress(&s, oid, ures.SecureURL, ures.TokenKey)
				if len(*sres) == 0 || (*sres)[0].Status != "201" {
					log.Println("Address Error: ", sres)
				}
				log.Println("Address: ", sres)
				session.Values["fflLic"] = nil
				session.Save(r, w)
				// var capi  api.API
				// capi.
				//log.Println("shipTo in ffl in address: ", res)
			} else {
				pg.NoFFL = true
			}
		} else {
			pg.NoFFL = true
		}
	} else {
		pg.NoFFL = true
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
