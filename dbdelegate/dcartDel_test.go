package dbdelegate

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mydb "github.com/Ulbora/dbinterface_mysql"
)

func TestDCartDeligate_init(t *testing.T) {

	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{"1"}
	mdb.MockTestRow = &mTestRow
	mdb.MockConnectSuccess = true

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	//dcart = dcDel.GetNew()
	suc := dcDel.DB.Connect()
	if !suc {
		t.Fail()
	}
}

func TestDCartDeligate_testConnection(t *testing.T) {

	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{"1"}
	mdb.MockTestRow = &mTestRow
	mdb.MockConnectSuccess = true

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	res := dcDel.testConnection()
	if !res {
		t.Fail()
	}
}

func TestDCartDeligate_AddUser(t *testing.T) {
	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{"1"}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "some user", "456211111", "somedomain.com", "customer", "1"}
	mdb.MockRow1 = &getRow

	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l

	dcart := dcDel.GetNew()
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	fmt.Println("dcart: ", dcart)
	suc, id := dcart.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	}
}

func TestDCartDeligate_AddUserNoRec(t *testing.T) {
	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{"1"}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockInsertSuccess1 = true
	mdb.MockInsertID1 = 5

	var getRow db.DbRow
	//getRow.Row = []string{"1", "some user", "456211111", "somedomain.com", "customer", "1"}
	getRow.Row = []string{}
	mdb.MockRow1 = &getRow

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	dcart := dcDel.GetNew()
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	fmt.Println("dcart: ", dcart)
	suc, id := dcart.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	}
}

func TestDCartDeligate_AddUserConnection(t *testing.T) {
	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "some user", "456211111", "somedomain.com", "customer", "1"}
	mdb.MockRow1 = &getRow

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	dcart := dcDel.GetNew()
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	fmt.Println("dcart: ", dcart)
	suc, id := dcart.AddUser(&dcu)
	if !suc || id < 1 {
		t.Fail()
	}
}

func TestDCartDeligate_RemoveUser(t *testing.T) {
	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{"1"}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "some user", "456211111", "somedomain.com", "customer", "1"}
	mdb.MockRow1 = &getRow

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	dcart := dcDel.GetNew()

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

func TestDCartDeligate_RemoveUserConnect(t *testing.T) {
	var dcDel DCartDeligate

	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "some user", "456211111", "somedomain.com", "customer", "1"}
	mdb.MockRow1 = &getRow

	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	dcart := dcDel.GetNew()

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

func TestDCartDeligate_GetUser(t *testing.T) {
	var dcDel DCartDeligate

	//var insrtId1 int64
	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{"1"}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "12345", "456211111", "sometoken", "AUTHORIZE", "", "", "1"}
	mdb.MockRow1 = &getRow

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	dcart := dcDel.GetNew()
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	fmt.Println("dcart: ", dcart)

	user := dcart.GetUser("https://teststore.cdcart.com")
	fmt.Println("found user: ", user)
	if !user.Enabled || user.Action != "AUTHORIZE" {
		t.Fail()
	}
}

func TestDCartDeligate_GetUserConnection(t *testing.T) {
	var dcDel DCartDeligate

	//var insrtId1 int64
	var mdb mydb.MyDBMock
	mdb.Host = "localhost:3306"
	mdb.User = "admin"
	mdb.Password = "admin"
	mdb.Database = "dcart_flic"

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mdb.MockTestRow = &mTestRow

	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var getRow db.DbRow
	getRow.Row = []string{"1", "12345", "456211111", "sometoken", "AUTHORIZE", "", "", "1"}
	mdb.MockRow1 = &getRow

	//dbi = &mdb
	dcDel.DB = &mdb
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	dcDel.Log = &l
	dcart := dcDel.GetNew()
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	fmt.Println("dcart: ", dcart)

	user := dcart.GetUser("https://teststore.cdcart.com")
	fmt.Println("found user: ", user)
	if !user.Enabled || user.Action != "AUTHORIZE" {
		t.Fail()
	}
}

func Test_clearURL(t *testing.T) {
	s := cleanURL("https://teststore.cdcart.com")
	if s != "teststore.cdcart.com" {
		t.Fail()
	}
}

func Test_clearURL2(t *testing.T) {
	s := cleanURL("http://teststore.cdcart.com")
	if s != "teststore.cdcart.com" {
		t.Fail()
	}
}
