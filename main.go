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
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	dcd "github.com/Ulbora/FindByFlic/dbdelegate"
	flc "github.com/Ulbora/FindByFlic/flicfinder"
	hand "github.com/Ulbora/FindByFlic/handlers"
	mydb "github.com/Ulbora/dbinterface_mysql"
	api "github.com/Ulbora/dcartapi"
	usession "github.com/Ulbora/go-better-sessions"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	//px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

const (
	userSession       = "ffl-user-session"
	sessingTimeToLive = (15 * 60) //120 minutes -- 2 hours
)

var templates *template.Template
var fh hand.FlicHandler
var s usession.Session

func main() {
	var l lg.Logger
	l.LogLevel = lg.AllLevel

	var privateKey string
	if len(os.Args) >= 2 {
		privateKey = os.Args[1]
	}
	if privateKey == "" {
		privateKey = os.Getenv("PRIVATE_KEY")
	}
	var dapi api.API
	dapi.PrivateKey = privateKey
	fh.DcartAPI = &dapi

	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
	} else {
		s.SessionKey = "115722gggg14ddfg4567"
	}
	fh.Sess = s

	userDb := connectUserDB(&fh)
	defer userDb.DB.Close()
	fmt.Println("del db: ", fh.DCartUserDel)

	//fflDb := connectFFLDB(&h)
	//defer fflDb.DB.Close()
	var flicFinder flc.FlicFinder

	flicFinder.Log = &l
	fh.FlicFinder = &flicFinder
	userDb.Log = &l

	//var gp px.GoProxy

	fmt.Println("del db: ", fh.FlicFinder)

	fh.Templates = template.Must(template.ParseFiles("./static/index.html", "./static/dcartIndex.html",
		"./static/dcartConfig.html", "./static/head.html", "./static/dcartAddFfl.html",
		"./static/dcartChosenFfl.html", "./static/dcartShippedFfl.html", "./static/dcartShipFflAddress.html"))
	//h.Templates = template.Must(template.ParseFiles("./static/index.html"))
	router := mux.NewRouter()
	h := fh.GetNew()
	//dcart
	router.HandleFunc("/", h.HandleIndex).Methods("GET")
	router.HandleFunc("/dcart", h.HandleDcartIndex).Methods("GET")
	router.HandleFunc("/dcartconfig", h.HandleDcartConfig).Methods("GET")
	router.HandleFunc("/dcartcb", h.HandleDcartCb).Methods("POST")
	router.HandleFunc("/dcartFindffl", h.HandleDcartFindFFL).Methods("POST")
	router.HandleFunc("/dcartChooseFFL", h.HandleDcartChooseFFL).Methods("GET")
	router.HandleFunc("/dcartShipffl", h.HandleDcartShipFFL).Methods("POST")
	router.HandleFunc("/dcartShipfflAddress", h.HandleDcartShipFFLAddress).Methods("GET")
	//dcart rest services
	router.HandleFunc("/dcart/rs/ffllist/{zip}", h.HandleFFLList).Methods("GET")
	router.HandleFunc("/dcart/rs/ffl/{id}", h.HandleFFLGet).Methods("GET")
	router.HandleFunc("/dcart/rs/shipment", h.HandleFFLAddAddress).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	//<iframe src="https://localhost:8060"></iframe>

	log.Println("Online Account Creator!")
	log.Println("Listening on :8070...")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "SecureURL", "Content-Type", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8070", handlers.CORS(headersOk, originsOk, methodsOk)(router))
	//http.ListenAndServeTLS(":8070", "certLocal.pem", "keyLocal.pem", router)
	//http.ListenAndServeTLS(":8070", "cert.pem", "key.pem", router)

}

func connectUserDB(h *hand.FlicHandler) *dcd.DCartDeligate {
	var dcDel dcd.DCartDeligate
	var mdb = new(mydb.MyDB)
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
	dcDel.DB = mdb
	suc := dcDel.DB.Connect()
	if !suc {
		log.Println("User DB connect failed")
	}
	h.DCartUserDel = &dcDel
	return &dcDel
}

// func connectFFLDB(h *hand.Handler) *ffl.Finder {
// 	var finder ffl.Finder
// 	var fflmdb = new(mydb.MyDB)
// 	if os.Getenv("FFL_DATABASE_HOST") != "" {
// 		fflmdb.Host = os.Getenv("FFL_DATABASE_HOST")
// 	} else {
// 		fflmdb.Host = "localhost:3306"
// 	}

// 	if os.Getenv("FFL_DATABASE_USER_NAME") != "" {
// 		fflmdb.User = os.Getenv("FFL_DATABASE_USER_NAME")
// 	} else {
// 		fflmdb.User = "admin"
// 	}

// 	if os.Getenv("FFL_DATABASE_USER_PASSWORD") != "" {
// 		fflmdb.Password = os.Getenv("FFL_DATABASE_USER_PASSWORD")
// 	} else {
// 		fflmdb.Password = "admin"
// 	}

// 	if os.Getenv("FFL_DATABASE_NAME") != "" {
// 		fflmdb.Database = os.Getenv("FFL_DATABASE_NAME")
// 	} else {
// 		fflmdb.Database = "ffl_list_10012018"
// 	}
// 	finder.DB = fflmdb
// 	suc := finder.DB.Connect()
// 	if !suc {
// 		log.Println("FFL DB connect failed")
// 	}
// 	h.FFLFinder = &finder
// 	return &finder
// }

//go mod init github.com/Ulbora/FindByFlic
