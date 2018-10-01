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

package dbdelegate

import (
	"fmt"
	dbi "github.com/Ulbora/dbinterface"
	"log"
	"strconv"
	"strings"
	"time"
)

//DCartUser DCartUser
type DCartUser struct {
	PublicKey string `json:"PublicKey"`
	TimeStamp string `json:"TimeStamp"`
	TokenKey  string `json:"TokenKey"`
	Action    string `json:"Action"`
	SecureURL string `json:"SecureURL"`
	Enabled   bool
}

//FindFFLDCart FindFFLDCart
type FindFFLDCart interface {
	AddUser(cu *DCartUser) (bool, int64)
	RemoveUser(cu *DCartUser) bool
	GetUser(url string) *DCartUser
}

//DCartDeligate DCartDeligate
type DCartDeligate struct {
	DB dbi.Database
}

//AddUser AddUser
func (d *DCartDeligate) AddUser(cu *DCartUser) (bool, int64) {
	var suc bool
	var id int64
	if cu.Action == "AUTHORIZE" {
		if !d.testConnection() {
			d.DB.Connect()
		}
		cu.SecureURL = cleanURL(cu.SecureURL)
		//fmt.Println("cu.SecureURL", cu.SecureURL)
		var a []interface{}
		a = append(a, cu.SecureURL)
		rowPtr := d.DB.Get(dcartGetByStore, a...)
		if rowPtr != nil {
			foundRow := rowPtr.Row
			if len(foundRow) > 0 {
				log.Println("Found existing record")
				fid, err := strconv.ParseInt(foundRow[0], 10, 64)
				if err != nil {
					log.Println("error converting id to int64: ", err)
				}
				var au []interface{}
				au = append(au, cu.PublicKey, cu.TokenKey, cu.Action, time.Now(), fid)
				usec := d.DB.Update(dcartUpdateStore, au...)
				suc = usec
				id = fid
			} else {
				log.Println("No record found inserting new record record")
				var au []interface{}
				au = append(au, cu.SecureURL, cu.PublicKey, cu.TokenKey, cu.Action, time.Now(), true)
				isuc, iid := d.DB.Insert(dcartInsertStore, au...)
				suc = isuc
				id = iid
			}
		}
	}
	return suc, id
}

//RemoveUser RemoveUser
func (d *DCartDeligate) RemoveUser(cu *DCartUser) bool {
	var rtn = false
	if cu.Action == "REMOVE" {
		if !d.testConnection() {
			d.DB.Connect()
		}
		cu.SecureURL = cleanURL(cu.SecureURL)
		var a []interface{}
		a = append(a, cu.SecureURL)
		rowPtr := d.DB.Get(dcartGetByStore, a...)
		if rowPtr != nil {
			foundRow := rowPtr.Row
			if len(foundRow) > 0 {
				log.Println("Found existing record")
				fid, err := strconv.ParseInt(foundRow[0], 10, 64)
				if err != nil {
					log.Println("error converting id to int64: ", err)
				}
				var au []interface{}
				au = append(au, cu.Action, time.Now(), fid)
				usec := d.DB.Update(dcartRemoveStore, au...)
				rtn = usec
			}
		}
	}
	return rtn
}

//GetUser GetUser
func (d *DCartDeligate) GetUser(url string) *DCartUser {
	if !d.testConnection() {
		log.Println("test database failed, reconnection database")
		d.DB.Connect()
	}
	var rtn DCartUser
	url = cleanURL(url)
	var a []interface{}
	a = append(a, url)
	rowPtr := d.DB.Get(dcartGetByStore, a...)
	if rowPtr != nil {
		//fid, err := strconv.ParseInt(foundRow[0], 10, 64)
		foundRow := rowPtr.Row
		rtn.SecureURL = foundRow[1]
		rtn.PublicKey = foundRow[2]
		rtn.TokenKey = foundRow[3]
		rtn.Action = foundRow[4]
		enabled, err := strconv.ParseBool(foundRow[7])
		if err != nil {
			log.Print(err)
		} else {
			rtn.Enabled = enabled
		}
		// cTime, errCtime := time.Parse(timeFormat, (*foundRow)[3])
		// if errCtime != nil {
		// 	fmt.Print(errCtime)
		// } else {
		// 	rtn.CreateDate = cTime
		// }

	}
	return &rtn
}

func (d *DCartDeligate) testConnection() bool {
	var rtn = false
	var a []interface{}
	rowPtr := d.DB.Test(dcartTest, a...)
	log.Println("rowPtr", rowPtr)
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		log.Print("Records found during test ")
		log.Println("Records found during test :", int64Val)
		if err != nil {
			log.Print(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

func cleanURL(url string) string {
	var rtn string
	if strings.Contains(url, "https:") {
		rtn = strings.TrimPrefix(url, "https://")
	} else {
		rtn = strings.TrimPrefix(url, "http://")
	}
	fmt.Println("url:", rtn)
	return rtn
}
