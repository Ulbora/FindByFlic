package fflfinder

//FFLFinder FFLFinder
type FFLFinder interface {
	FindFFL(zip string) *[]FFL
}

//FFL FFL
type FFL struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	LicNumber string `json:"licNumber"`
}
