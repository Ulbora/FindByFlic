package dbdelegate

import (
	"fmt"
	"testing"

	dbi "github.com/Ulbora/dbinterface"
	mydb "github.com/Ulbora/dbinterface/mysql"
)

var dcart FindFFLDCart
var dcDel DCartDeligate
var db dbi.Database

var insrtId1 int64

func TestDCartDeligate_init(t *testing.T) {
	var mdb mydb.MyDB
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"
	db = &mdb
	dcDel.DB = db
	dcart = &dcDel
	suc := dcDel.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestDCartDeligate_testConnection(t *testing.T) {
	res := dcDel.testConnection()
	if !res {
		t.Fail()
	}
}

func TestDCartDeligate_AddUser(t *testing.T) {
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	suc, id := dcart.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	} else {
		insrtId1 = id
	}
}

func TestDCartDeligate_RemoveUser(t *testing.T) {
	var dcu DCartUser
	dcu.Action = "REMOVE"
	//dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	//dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" //time.Now()
	suc := dcart.RemoveUser(&dcu)
	if !suc {
		t.Fail()
	}
}

func TestDCartDeligate_AddUser2(t *testing.T) {
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" //time.Now()
	suc, id := dcart.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	}
}

func TestDCartDeligate_GetUser(t *testing.T) {
	user := dcart.GetUser("https://teststore.cdcart.com")
	fmt.Println("found user: ", user)
	if !user.Enabled || user.Action != "AUTHORIZE" {
		t.Fail()
	}
}

func TestDCartDeligate_delete(t *testing.T) {
	var a []interface{}
	a = append(a, insrtId1)
	suc := dcDel.DB.Delete("delete from dcart_user where id = ?", a...)

	if !suc {
		t.Fail()
	}
}

func TestDCartDeligate_close(t *testing.T) {
	suc := dcDel.DB.Close()
	fmt.Println("closing db")
	if !suc {
		t.Fail()
	}
}

func Test_clearURL(t *testing.T) {
	s := cleanURL("https://teststore.cdcart.com")
	if s != "teststore.cdcart.com" {
		t.Fail()
	}
}
