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

package fflfinder

import (
	dbi "github.com/Ulbora/dbinterface"
	"log"
	"strconv"
)

//Finder Finder
type Finder struct {
	DB dbi.Database
}

// SELECT *
// FROM `ffl_lic`
// WHERE premise_zip_code like '30132%'
// and ((lic_type = '07') xor (lic_type = '01') or lic_type = '02')

// SELECT *
// FROM `ffl_lic`
// WHERE premise_zip_code like '30132%'
// and ((lic_type = '07') xor (lic_type = '01') or lic_type = '02')
// order by license_name ASC

//FindFFL FindFFL
func (f *Finder) FindFFL(zip string) *[]FFL {
	var rtn []FFL
	if !f.testConnection() {
		f.DB.Connect()
	}
	var a []interface{}
	a = append(a, zip+"%")
	rowPtr := f.DB.GetList(fflListQuery, a...)
	if len(rowPtr.Rows) > 0 {
		rows := rowPtr.Rows
		for _, r := range rows {
			//var ffl FFL
			// id, _ := strconv.ParseInt(r[0], 10, 64)
			// ffl.ID = id
			// ffl.LicRegn = r[1]
			// ffl.LicDist = r[2]
			// ffl.LicCnty = r[3]
			// ffl.LicType = r[4]
			// ffl.LicXprdte = r[5]
			// ffl.LicSeqn = r[6]
			// ffl.LicenseName = r[7]
			// ffl.BusinessName = r[8]
			// ffl.PremiseStreet = r[9]
			// ffl.PremiseCity = r[10]
			// ffl.PremiseState = r[11]
			// ffl.PremiseZipCode = r[12]
			// ffl.VoicePhone = r[17]
			ffl := f.processFFL(&r)
			rtn = append(rtn, *ffl)
		}
	}
	return &rtn
}

//GetFFL GetFFL
func (f *Finder) GetFFL(id int64) *FFL {
	var rtn = new(FFL)
	if !f.testConnection() {
		f.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	rowPtr := f.DB.Get(fflGetQueru, a...)
	if len(rowPtr.Row) > 0 {
		rtn = f.processFFL(&rowPtr.Row)
	}
	return rtn
}

func (f *Finder) testConnection() bool {
	var rtn = false
	var a []interface{}
	rowPtr := f.DB.Test(fflTest, a...)
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

func (f *Finder) processFFL(r *[]string) *FFL {
	var ffl FFL
	id, _ := strconv.ParseInt((*r)[0], 10, 64)
	ffl.ID = id
	ffl.LicRegn = (*r)[1]
	ffl.LicDist = (*r)[2]
	ffl.LicCnty = (*r)[3]
	ffl.LicType = (*r)[4]
	ffl.LicXprdte = (*r)[5]
	ffl.LicSeqn = (*r)[6]
	ffl.LicenseName = (*r)[7]
	ffl.BusinessName = (*r)[8]
	ffl.PremiseStreet = (*r)[9]
	ffl.PremiseCity = (*r)[10]
	ffl.PremiseState = (*r)[11]
	ffl.PremiseZipCode = (*r)[12]
	ffl.VoicePhone = (*r)[17]
	return &ffl
}
