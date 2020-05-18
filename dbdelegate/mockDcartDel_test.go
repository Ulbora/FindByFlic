package dbdelegate

import (
	"testing"
)

func TestMockDCartDeligate_AddUser(t *testing.T) {
	var del MockDCartDeligate
	del.MockAddID = 5
	del.MockAddSuccess = true
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	d := del.GetNew()
	suc, id := d.AddUser(&dcu)
	if !suc || id == 0 {
		t.Fail()
	}
}

func TestMockDCartDeligate_RemoveUser(t *testing.T) {
	var del MockDCartDeligate
	//del.MockAddID = 5
	del.MockRemoveSuccess = true
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	d := del.GetNew()
	suc := d.RemoveUser(&dcu)
	if !suc {
		t.Fail()
	}
}

func TestMockDCartDeligate_GetUser(t *testing.T) {
	var del MockDCartDeligate
	//del.MockAddID = 5
	//del.MockRemoveSuccess = true
	var dcu DCartUser
	dcu.Action = "AUTHORIZE"
	dcu.PublicKey = "12345"
	dcu.SecureURL = "https://teststore.cdcart.com"
	dcu.TokenKey = "555ggg11"
	dcu.TimeStamp = "12-25-2018 01:01:00" // time.Now()
	dcu.Enabled = true
	del.MockDcartUser = &dcu
	d := del.GetNew()
	user := d.GetUser("https://teststore.cdcart.com")
	if !user.Enabled {
		t.Fail()
	}
}
