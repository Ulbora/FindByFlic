package dbdeligate

import (
	dbi "github.com/Ulbora/dbinterface"
	"log"
	"strconv"
	"time"
)

//DCartUser DCartUser
type DCartUser struct {
	PublicKey string    `json:"PublicKey"`
	TimeStamp time.Time `json:"TimeStamp"`
	TokenKey  string    `json:"TokenKey"`
	Action    string    `json:"Action"`
	SecureURL string    `json:"SecureURL"`
}

//FindFFLDCart FindFFLDCart
type FindFFLDCart interface {
	AddUser(cu DCartUser) (bool, int64)
	RemoveUser(cu DCartUser) bool
	GetUser(url string) *dbi.DbRow
}

//DCartDeligate DCartDeligate
type DCartDeligate struct {
	DB dbi.Database
}

//AddUser AddUser
func (d *DCartDeligate) AddUser(cu DCartUser) (bool, int64) {
	var suc bool
	var id int64
	if cu.Action == "AUTHORIZE" {

	}
	return suc, id
}

//RemoveUser RemoveUser
func (d *DCartDeligate) RemoveUser(cu DCartUser) bool {
	return false
}

//GetUser GetUser
func (d *DCartDeligate) GetUser(url string) *dbi.DbRow {
	return new(dbi.DbRow)
}

func (d *DCartDeligate) testConnection() bool {
	var rtn = false
	var a []interface{}
	rowPtr := d.DB.Test(dcartTest, a...)
	if rowPtr != nil {
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
