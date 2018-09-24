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

package main

import (
	dcd "FindByFlic/dbdelegate"
	hand "FindByFlic/handlers"
	"fmt"
	dbi "github.com/Ulbora/dbinterface"
	mydb "github.com/Ulbora/dbinterface/mysql"
	usession "github.com/Ulbora/go-better-sessions"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	userSession       = "ffl-user-session"
	sessingTimeToLive = (15 * 60) //120 minutes -- 2 hours
)

var templates *template.Template
var h hand.Handler
var s usession.Session

var dcart dcd.FindFFLDCart
var dcDel dcd.DCartDeligate
var db dbi.Database
var mdb mydb.MyDB

func main() {

	if os.Getenv("DATABASE_HOST") != "" {
		mdb.Host = os.Getenv("DATABASE_HOST")
	} else {
		mdb.Host = "localhost:3306"
	}

	if os.Getenv("DATABASE_USER_NAME") != "" {
		mdb.User = os.Getenv("DATABASE_USER_NAME")
	} else {
		mdb.User = "admin"
	}

	if os.Getenv("DATABASE_USER_PASSWORD") != "" {
		mdb.Password = os.Getenv("DATABASE_USER_PASSWORD")
	} else {
		mdb.Password = "admin"
	}

	if os.Getenv("DATABASE_NAME") != "" {
		mdb.Database = os.Getenv("DATABASE_NAME")
	} else {
		mdb.Database = "dcart_flic"
	}
	db = &mdb
	dcDel.DB = db
	dcart = &dcDel
	suc := dcDel.DB.Connect()
	if suc {
		defer dcDel.DB.Close()
		h.FindFFLDCart = dcart
		fmt.Println("del db: ", h.FindFFLDCart)

		s.MaxAge = sessingTimeToLive
		s.Name = userSession
		if os.Getenv("SESSION_SECRET_KEY") != "" {
			s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
		} else {
			s.SessionKey = "115722gggg14ddfg4567"
		}
		h.Sess = s
		h.Templates = template.Must(template.ParseFiles("./static/index.html", "./static/dcartIndex.html",
			"./static/dcartConfig.html", "./static/head.html"))
		//h.Templates = template.Must(template.ParseFiles("./static/index.html"))
		router := mux.NewRouter()
		router.HandleFunc("/", h.HandleIndex).Methods("GET")
		router.HandleFunc("/dcart", h.HandleDcartIndex).Methods("GET")
		router.HandleFunc("/dcartconfig", h.HandleDcartConfig).Methods("GET")
		router.HandleFunc("/dcartcb", h.HandleDcartCb).Methods("POST")

		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

		//<iframe src="https://localhost:8060"></iframe>

		log.Println("Online Account Creator!")
		log.Println("Listening on :8070...")
		http.ListenAndServe(":8070", router)
		//http.ListenAndServeTLS(":8070", "certLocal.pem", "keyLocal.pem", router)
		//http.ListenAndServeTLS(":8070", "cert.pem", "key.pem", router)

	}

}
