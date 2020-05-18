package flicfinder

//Finder Finder
type Finder interface {
	FindFlicListByZip(zip string) *[]FlicList
	FindFlicByID(id string) *Flic
}

//FlicList FlicList
type FlicList struct {
	Key            string `json:"id"`
	LicName        string `json:"licenseName"`
	BusName        string `json:"businessName"`
	PremiseAddress string `json:"premiseAddress"`
}

//Flic Flic
type Flic struct {
	Key            string `json:"id"`
	Lic            string `json:"license"`
	ExpDate        string `json:"expDate"`
	LicName        string `json:"licenseName"`
	BusName        string `json:"businessName"`
	PremiseAddress string `json:"premiseAddress"`
	Address        string `json:"address"`
	City           string `json:"city"`
	State          string `json:"state"`
	PremiseZip     string `json:"premiseZip"`
	MailingAddress string `json:"mailingAddress"`
	Phone          string `json:"phone"`
}
