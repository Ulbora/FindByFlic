package handlers

import (
	"encoding/json"
	"fmt"
	api "github.com/Ulbora/dcartapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"

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
func (h *Handler) HandleFFLList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := r.Header.Get("SecureURL")
	ures := h.FindFFLDCart.GetUser(url)
	//log.Println("user found before if: ", ures)
	if ures.Enabled {
		var zip string
		vars := mux.Vars(r)
		if vars != nil {
			zip = vars["zip"]
		} else {
			zip = r.URL.Query().Get("zip")
		}
		res := h.FFLFinder.FindFFL(zip)
		//log.Println("rest ffl list: ", res)
		resJSON, err := json.Marshal(res)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "json output failed", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//HandleFFLGet HandleFFLGet
func (h *Handler) HandleFFLGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.Header.Get("SecureURL")
	ures := h.FindFFLDCart.GetUser(url)
	//log.Println("user found before if: ", ures)
	if ures.Enabled {
		var idstr string
		vars := mux.Vars(r)
		if vars != nil {
			idstr = vars["id"]
		} else {
			idstr = r.URL.Query().Get("id")
		}
		id, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			res := h.FFLFinder.GetFFL(id)
			//log.Println("rest ffl: ", res)
			resJSON, err := json.Marshal(res)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//HandleFFLAddAddress HandleFFLAddAddress
func (h *Handler) HandleFFLAddAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		url := r.Header.Get("SecureURL")
		ures := h.FindFFLDCart.GetUser(url)
		log.Println("user in rest: ", ures)
		if ures.Enabled {
			//log.Println("req: ", r.Body)
			addReq := new(AddressRequest)
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(addReq)
			if error != nil {
				log.Println(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			} else if addReq.Invoice == "" || addReq.FFLID == "" {
				http.Error(w, "bad request", http.StatusBadRequest)
			} else {
				//log.Println("ffl lic in address: ", addReq.FFLID)
				id, err := strconv.ParseInt(addReq.FFLID, 10, 64)
				if err != nil {
					log.Println("ffl ID error: ", err)
				}
				res := h.FFLFinder.GetFFL(id)
				log.Println("ffl lic rest: ", res)
				log.Println("addReq before getOrder: ", addReq)
				odr := h.DcartAPI.GetOrder(addReq.Invoice, ures.SecureURL, ures.TokenKey)
				log.Println("odr: ", odr)
				if odr.OrderID != 0 && res.ID != 0 {
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
					var rtn AddressResponse
					sres := h.DcartAPI.AddShippingAddress(&s, oid, ures.SecureURL, ures.TokenKey)
					if len(*sres) == 0 || (*sres)[0].Status != "201" {
						log.Println("Address Error: ", sres)
						w.WriteHeader(http.StatusBadRequest)
					} else {
						rtn.Success = true
						w.WriteHeader(http.StatusOK)
					}
					resJSON, err := json.Marshal(rtn)
					//log.Println("rtn rest: ", rtn)
					if err != nil {
						log.Println(error.Error())
						http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					fmt.Fprint(w, string(resJSON))
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
