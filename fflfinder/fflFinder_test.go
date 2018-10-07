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
	"fmt"

	"testing"

	dbi "github.com/Ulbora/dbinterface"
	mydb "github.com/Ulbora/dbinterface/mysql"
)

var fflFinder FFLFinder
var finder Finder
var db dbi.Database

func TestFinder_init(t *testing.T) {
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "ffl_list_10012018"
	db = &mdb
	finder.DB = db
	fflFinder = &finder
	suc := finder.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestFinder_testConnection(t *testing.T) {
	res := finder.testConnection()
	if !res {
		t.Fail()
	}
}

func TestFinder_FindFFL(t *testing.T) {
	res := fflFinder.FindFFL("30144")
	fmt.Println("ffl list: ", res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestFinder_GetFFL(t *testing.T) {
	res := fflFinder.GetFFL(5)
	fmt.Println("ffl 5: ", res)
	if res.ID == 0 {
		t.Fail()
	}
}

func TestFinder_close(t *testing.T) {
	suc := finder.DB.Close()
	fmt.Println("closing db")
	if !suc {
		t.Fail()
	}
}
