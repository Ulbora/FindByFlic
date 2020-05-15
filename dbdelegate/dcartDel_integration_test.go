// +build integration move to top

package dbdelegate

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mydb "github.com/Ulbora/dbinterface_mysql"
)

var dcarti FindFFLDCart
var dcDeli DCartDeligate

//var dbi dbi.Database

var insrtId1i int64

func TestDCartDeligatei_init(t *testing.T) {
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"
	//dbi = &mdb
	dcDeli.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDeli.Log = &l
	dcarti = dcDeli.GetNew()
	suc := dcDeli.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestDCartDeligatei_testConnection(t *testing.T) {
	res := dcDeli.testConnection()
	if !res {
		t.Fail()
	}
}

func TestDCartDeligatei_AddUser(t *testing.T) {
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	suc, id := dcarti.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	} else {
		insrtId1i = id
	}
}

func TestDCartDeligatei_RemoveUser(t *testing.T) {
	var dcu DCartUser
	dcu.Action = "REMOVE"
	//dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	//dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" //time.Now()
	suc := dcarti.RemoveUser(&dcu)
	if !suc {
		t.Fail()
	}
}

func TestDCartDeligatei_AddUser2(t *testing.T) {
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" //time.Now()
	suc, id := dcarti.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	}
}

func TestDCartDeligatei_GetUser(t *testing.T) {
	user := dcarti.GetUser("https://teststore.cdcart.com")
	fmt.Println("found user: ", user)
	if !user.Enabled || user.Action != "AUTHORIZE" {
		t.Fail()
	}
}

func TestDCartDeligatei_delete(t *testing.T) {
	var a []interface{}
	a = append(a, insrtId1i)
	suc := dcDeli.DB.Delete("delete from dcart_user where id = ?", a...)

	if !suc {
		t.Fail()
	}
}

func TestDCartDeligatei_close(t *testing.T) {
	suc := dcDeli.DB.Close()
	fmt.Println("closing db")
	if !suc {
		t.Fail()
	}
}

func Test_clearURLi(t *testing.T) {
	s := cleanURL("https://teststore.cdcart.com")
	if s != "teststore.cdcart.com" {
		t.Fail()
	}
}

func Test_clearURL2i(t *testing.T) {
	s := cleanURL("http://teststore.cdcart.com")
	if s != "teststore.cdcart.com" {
		t.Fail()
	}
}
