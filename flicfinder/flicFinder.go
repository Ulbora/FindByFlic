package flicfinder

import (
	"bytes"
	"encoding/json"
	"net/http"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

//FlicFinder FlicFinder
type FlicFinder struct {
	Proxy   px.Proxy
	APIKey  string
	FlicURL string
	Log     *lg.Logger
}

//FlicRequest FlicRequest
type FlicRequest struct {
	Zip string `json:"zip"`
	ID  string `json:"id"`
}

//GetNew GetNew
func (f *FlicFinder) GetNew() Finder {
	return f
}

//FindFlicListByZip FindFlicListByZip
func (f *FlicFinder) FindFlicListByZip(zip string) *[]FlicList {
	var r FlicRequest
	r.Zip = zip
	var rtn = []FlicList{}
	aJSON, _ := json.Marshal(r)
	req, rErr := http.NewRequest("POST", f.FlicURL+"/rs/findByZip", bytes.NewBuffer(aJSON))
	f.Log.Debug("request err:  ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("api-key", f.APIKey)
		f.Log.Debug("m.Proxy:  ", f.Proxy)
		suc, code := f.Proxy.Do(req, &rtn)
		if !suc || code != http.StatusOK {
			f.Log.Debug("suc:  ", suc)
			f.Log.Debug("code:  ", code)
		}
	}
	return &rtn
}

//FindFlicByID FindFlicByID
func (f *FlicFinder) FindFlicByID(id string) *Flic {
	var r FlicRequest
	r.ID = id
	var rtn Flic
	aJSON, _ := json.Marshal(r)
	req, rErr := http.NewRequest("POST", f.FlicURL+"/rs/findById", bytes.NewBuffer(aJSON))
	f.Log.Debug("request err:  ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("api-key", f.APIKey)
		f.Log.Debug("m.Proxy:  ", f.Proxy)
		suc, code := f.Proxy.Do(req, &rtn)
		if !suc || code != http.StatusOK {
			f.Log.Debug("suc:  ", suc)
			f.Log.Debug("code:  ", code)
		}
	}
	return &rtn
}
