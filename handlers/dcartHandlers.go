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
	"strconv"

	dcd "github.com/Ulbora/FindByFlic/dbdelegate"
	flc "github.com/Ulbora/FindByFlic/flicfinder"
	api "github.com/Ulbora/dcartapi"

	//dbi "github.com/Ulbora/dbinterface"
	//ffl "FindByFlic/fflfinder"

	"net/http"
)

//HandleDcartIndex HandleDcartIndex
func (h *FlicHandler) HandleDcartIndex(w http.ResponseWriter, r *http.Request) {
	//h.Sess.InitSessionStore(w, r)
	var carturl = r.URL.Query().Get("carturl")
	h.Log.Debug("carturl in ffl address: ", carturl)
	ures := h.DCartUserDel.GetUser(carturl)
	h.Log.Debug("ures in index: ", ures)
	var pg PageParams
	if ures != nil && ures.Enabled {
		pg.Enabled = "true"
	} else {
		pg.Enabled = "false"
	}
	h.Templates.ExecuteTemplate(w, "dcartIndex.html", &pg)
}

//HandleDcartConfig HandleDcartConfig
func (h *FlicHandler) HandleDcartConfig(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	h.Templates.ExecuteTemplate(w, "dcartConfig.html", nil)
}

//HandleDcartCb HandleDcartCb
func (h *FlicHandler) HandleDcartCb(w http.ResponseWriter, r *http.Request) {
	var rtn bool
	h.Log.Debug("service call from 3dcart on callback-------")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		dcReg := new(dcd.DCartUser)
		decoder := json.NewDecoder(r.Body)
		error := decoder.Decode(dcReg)
		if error != nil {
			h.Log.Debug(error.Error())
			http.Error(w, error.Error(), http.StatusBadRequest)
		} else {
			if dcReg.Action == "AUTHORIZE" {
				rtn, _ = h.DCartUserDel.AddUser(dcReg)
			} else if dcReg.Action == "REMOVE" {
				rtn = h.DCartUserDel.RemoveUser(dcReg)
			}
			h.Log.Debug("add user to cart success: ", rtn)

			if rtn {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}

//HandleDcartFindFFL HandleDcartFindFFL
func (h *FlicHandler) HandleDcartFindFFL(w http.ResponseWriter, r *http.Request) {
	zip := r.FormValue("zip")
	use := r.FormValue("use")
	var res = new([]flc.FlicList)
	if use == "true" {
		//res = h.FFLFinder.FindFFL(zip)
		res = h.FlicFinder.FindFlicListByZip(zip)
	}
	var pg FFLPageParams
	pg.FFLList = res
	pg.Zip = zip
	pg.Enabled = use
	h.Log.Debug("ffl list in dcart handler: ", *pg.FFLList)
	h.Templates.ExecuteTemplate(w, "dcartAddFfl.html", &pg)
}

//HandleDcartChooseFFL HandleDcartChooseFFL
func (h *FlicHandler) HandleDcartChooseFFL(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	zip := r.FormValue("zip")
	use := r.FormValue("use")
	//id, _ := strconv.ParseInt(idstr, 10, 64)
	res := h.FlicFinder.FindFlicByID(id)
	var pg FFLPageParams
	pg.FFL = res
	pg.Zip = zip
	pg.Enabled = use
	h.Log.Debug("ffl in choose: ", pg.FFL)
	h.Templates.ExecuteTemplate(w, "dcartChosenFfl.html", &pg)
}

//HandleDcartShipFFL HandleDcartShipFFL
func (h *FlicHandler) HandleDcartShipFFL(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	id := r.FormValue("id")
	use := r.FormValue("use")
	//id, _ := strconv.ParseInt(idstr, 10, 64)
	//log.Println("licNum in ship: ", licNum)
	//name := r.FormValue("name")
	//log.Println("name in ship: ", name)
	//address := r.FormValue("address")
	// log.Println("address in ship: ", address)
	// log.Println("order: ", session.Values["order"])
	// log.Println("carturl: ", session.Values["carturl"])
	/////res := h.FFLFinder.GetFFL(id)
	res := h.FlicFinder.FindFlicByID(id)
	h.Log.Debug("ffl in HandleDcartShipFFL: ", *res)
	//var theFFL interface{}
	//theFFL = *res
	session.Values["fflLic"] = id
	//session.Values["fflName"] = name
	//session.Values["fflAddress"] = address
	serr := session.Save(r, w)
	h.Log.Debug("session save err in HandleDcartShipFFL: ", serr)
	// if serr != nil {
	// 	log.Println("Session Err:", serr)
	// }
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
func (h *FlicHandler) HandleDcartShipFFLAddress(w http.ResponseWriter, r *http.Request) {
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
	h.Log.Debug("invoice in ffl address: ", invoice)
	h.Log.Debug("carturl in ffl address: ", carturl)
	//log.Println("shipTo: ", session.Values["shipTo"])
	var pg FFLPageParams
	fflLicID := session.Values["fflLic"]
	h.Log.Debug("fflLic: ", fflLicID)
	//var licNum string
	if fflLicID != nil {
		//idstr := fflLicID.(string)
		id := fflLicID.(string)
		///id := int64(idint)
		h.Log.Debug("ffl id: ", id)
		//id, _ := strconv.ParseInt(idstr, 10, 64)
		ures := h.DCartUserDel.GetUser(carturl)
		h.Log.Debug("user found before if: ", ures)
		if ures.Enabled {
			h.Log.Debug("user found: ", ures)
			//licNum = fflLic.(string)
			h.Log.Debug("ffl lic in address: ", id)
			res := h.FlicFinder.FindFlicByID(id)
			pg.FFL = res

			odr := h.DcartAPI.GetOrder(invoice, ures.SecureURL, ures.TokenKey)
			if odr.OrderID != 0 {
				var s api.Shipment
				s.ShipmentID = 0
				s.ShipmentFirstName = "FFL"
				s.ShipmentLastName = "Lic # " + res.Lic
				if res.BusName != "" {
					s.ShipmentCompany = res.BusName
				} else {
					s.ShipmentCompany = res.LicName
				}

				s.ShipmentAddress = res.Address
				s.ShipmentCity = res.City
				s.ShipmentState = res.State
				s.ShipmentZipCode = res.PremiseZip
				s.ShipmentCountry = "USA"
				s.ShipmentPhone = res.Phone
				s.ShipmentTax = 0
				s.ShipmentWeight = 1
				s.ShipmentTrackingCode = ""
				s.ShipmentUserID = ""
				s.ShipmentNumber = 1
				s.ShipmentAddressTypeID = 0
				var oid = strconv.FormatInt(odr.OrderID, 10)
				sres := h.DcartAPI.AddShippingAddress(&s, oid, ures.SecureURL, ures.TokenKey)
				if len(*sres) == 0 || (*sres)[0].Status != "201" {
					h.Log.Debug("Address Error: ", sres)
				}
				h.Log.Debug("Address: ", sres)
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
