package dbdelegate

import (
	lg "github.com/Ulbora/Level_Logger"
	dbi "github.com/Ulbora/dbinterface"
)

//MockDCartDeligate MockDCartDeligate
type MockDCartDeligate struct {
	DB                dbi.Database
	Log               *lg.Logger
	MockAddSuccess    bool
	MockAddID         int64
	MockRemoveSuccess bool
	MockDcartUser     *DCartUser
}

//GetNew GetNew
func (d *MockDCartDeligate) GetNew() DCartUserDelegate {
	return d
}

//AddUser AddUser
func (d *MockDCartDeligate) AddUser(cu *DCartUser) (bool, int64) {
	return d.MockAddSuccess, d.MockAddID
}

//RemoveUser RemoveUser
func (d *MockDCartDeligate) RemoveUser(cu *DCartUser) bool {
	return d.MockRemoveSuccess
}

//GetUser GetUser
func (d *MockDCartDeligate) GetUser(url string) *DCartUser {
	return d.MockDcartUser
}
