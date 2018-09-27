package fflfinder

//MockFinder MockFinder
type MockFinder struct {
}

//FindFFL FindFFL
func (f *MockFinder) FindFFL(zip string) *[]FFL {
	var rtn []FFL
	var ffl1 FFL
	ffl1.Name = "Wild West Traders"
	ffl1.Address = "123 Austell Rd, Marietta, GA 30166"
	ffl1.LicNumber = "456778888"
	rtn = append(rtn, ffl1)

	var ffl2 FFL
	ffl1.Name = "Wild Bills Traders"
	ffl1.Address = "123 Windy Hill Rd, Marietta, GA 30159"
	ffl1.LicNumber = "78899999"
	rtn = append(rtn, ffl2)

	var ffl3 FFL
	ffl1.Name = "Bill's Trading"
	ffl1.Address = "123 South Cobb Dr, Marietta, GA 30112"
	ffl1.LicNumber = "1123554"
	rtn = append(rtn, ffl3)

	var ffl4 FFL
	ffl1.Name = "Bills Guns"
	ffl1.Address = "123 Atlanta Rd, Marietta, GA 30166"
	ffl1.LicNumber = "7894562222"
	rtn = append(rtn, ffl4)
	return &rtn
}
