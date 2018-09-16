package dbdeligate

import (
	"testing"

	dbi "github.com/Ulbora/dbinterface"
	mydb "github.com/Ulbora/dbinterface/mysql"
)

var dcart FindFFLDCart

func TestDCartDeligate_init(t *testing.T) {
	var db dbi.Database
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "testdb"
	db = &mdb
	suc := db.Connect()
	if !suc {
		t.Fail()
	}
}

// func TestDCartDeligate_testConnection(t *testing.T) {

// }
