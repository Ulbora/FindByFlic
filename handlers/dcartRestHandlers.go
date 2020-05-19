package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "github.com/Ulbora/dcartapi"
	"github.com/gorilla/mux"

	"strconv"
)

//AddressRequest AddressRequest
type AddressRequest struct {
	Invoice string `json:"invoice"`
	FFLID   string `json:"fflId"`
}

//AddressResponse AddressResponse
type AddressResponse struct {
	Success bool `json:"success"`
}

//HandleFFLList HandleFFLList
func (h *FlicHandler) HandleFFLList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.Header.Get("SecureURL")
	ures := h.DCartUserDel.GetUser(url)
	h.Log.Debug("user found before if: ", *ures)
	if ures.Enabled {
		var zip string
		vars := mux.Vars(r)
		if vars != nil && len(vars) != 0 {
			zip = vars["zip"]
			h.Log.Debug("zip: ", zip)
			res := h.FlicFinder.FindFlicListByZip(zip)
			h.Log.Debug("rest ffl list: ", *res)
			resJSON, _ := json.Marshal(res)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//HandleFFLGet HandleFFLGet
func (h *FlicHandler) HandleFFLGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.Header.Get("SecureURL")
	ures := h.DCartUserDel.GetUser(url)
	h.Log.Debug("user found before if: ", *ures)
	if ures.Enabled {
		var id string
		vars := mux.Vars(r)
		if vars != nil && len(vars) != 0 {
			id = vars["id"]
			h.Log.Debug("id: ", id)
			res := h.FlicFinder.FindFlicByID(id)
			h.Log.Debug("rest ffl: ", *res)
			resJSON, _ := json.Marshal(res)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//HandleFFLAddAddress HandleFFLAddAddress
func (h *FlicHandler) HandleFFLAddAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		url := r.Header.Get("SecureURL")
		ures := h.DCartUserDel.GetUser(url)
		h.Log.Debug("user in rest: ", *ures)
		if ures.Enabled {
			addReq := new(AddressRequest)
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(addReq)
			if error != nil {
				h.Log.Debug(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			} else if addReq.Invoice == "" || addReq.FFLID == "" {
				http.Error(w, "bad request", http.StatusBadRequest)
			} else {
				res := h.FlicFinder.FindFlicByID(addReq.FFLID)
				h.Log.Debug("ffl lic rest: ", *res)
				h.Log.Debug("addReq before getOrder: ", *addReq)
				odr := h.DcartAPI.GetOrder(addReq.Invoice, ures.SecureURL, ures.TokenKey)
				h.Log.Debug("odr: ", *odr)
				if odr.OrderID != 0 && res.Key != "" {
					var s api.Shipment
					s.ShipmentID = 0
					s.ShipmentFirstName = "FFL"
					s.ShipmentLastName = "Lic # " + res.Lic
					if res.BusName != "" {
						s.ShipmentCompany = res.BusName
					} else {
						s.ShipmentCompany = res.LicName
					}
					s.ShipmentAddress = res.PremiseAddress
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
					var rtn AddressResponse
					sres := h.DcartAPI.AddShippingAddress(&s, oid, ures.SecureURL, ures.TokenKey)
					if len(*sres) == 0 || (*sres)[0].Status != "201" {
						h.Log.Debug("Address Error: ", sres)
						w.WriteHeader(http.StatusBadRequest)
					} else {
						rtn.Success = true
						w.WriteHeader(http.StatusOK)
						resJSON, _ := json.Marshal(rtn)
						fmt.Fprint(w, string(resJSON))
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
