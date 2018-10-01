/*
 Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2018 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
	ffl1.Address2 = "123 Austell Rd"
	ffl1.City = "Marietta"
	ffl1.State = "GA"
	ffl1.Zip = "30133"
	ffl1.Country = "USA"
	ffl1.Phone = "800 123 4567"
	ffl1.LicNumber = "456778888"
	rtn = append(rtn, ffl1)

	var ffl2 FFL
	ffl2.Name = "Wild Bills Traders"
	ffl2.Address = "123 Windy Hill Rd, Marietta, GA 30159"
	ffl2.Address2 = "123 Austell Rd"
	ffl2.City = "Marietta"
	ffl2.State = "GA"
	ffl2.Zip = "30133"
	ffl2.Country = "USA"
	ffl2.Phone = "800 123 4567"
	ffl2.LicNumber = "78899999"
	rtn = append(rtn, ffl2)

	var ffl3 FFL
	ffl3.Name = "Bill's Trading"
	ffl3.Address = "123 South Cobb Dr, Marietta, GA 30112"
	ffl3.Address2 = "123 Austell Rd"
	ffl3.City = "Marietta"
	ffl3.State = "GA"
	ffl3.Zip = "30133"
	ffl3.Country = "USA"
	ffl3.Phone = "800 123 4567"
	ffl3.LicNumber = "1123554"
	rtn = append(rtn, ffl3)

	var ffl4 FFL
	ffl4.Name = "Bills Guns"
	ffl4.Address = "123 Atlanta Rd, Marietta, GA 30166"
	ffl4.Address2 = "123 Austell Rd"
	ffl4.City = "Marietta"
	ffl4.State = "GA"
	ffl4.Zip = "30133"
	ffl4.Country = "USA"
	ffl4.Phone = "800 123 4567"
	ffl4.LicNumber = "7894562222"
	rtn = append(rtn, ffl4)

	var ffl5 FFL
	ffl5.Name = "Bills Arms"
	ffl5.Address = "123 Atlanta Rd, Actworth, GA 30166"
	ffl5.Address2 = "123 Austell Rd"
	ffl5.City = "Marietta"
	ffl5.State = "GA"
	ffl5.Zip = "30133"
	ffl5.Country = "USA"
	ffl5.Phone = "800 123 4567"
	ffl5.LicNumber = "7894562222"
	rtn = append(rtn, ffl5)

	var ffl6 FFL
	ffl6.Name = "Johns Guns"
	ffl6.Address = "123 Windy Hill Rd, Marietta, GA 30166"
	ffl6.Address2 = "123 Austell Rd"
	ffl6.City = "Marietta"
	ffl6.State = "GA"
	ffl6.Zip = "30133"
	ffl6.Country = "USA"
	ffl6.Phone = "800 123 4567"
	ffl6.LicNumber = "7894562222"
	rtn = append(rtn, ffl6)
	return &rtn
}

//GetFFL GetFFL
func (f *MockFinder) GetFFL(licNum string) *FFL {
	var ffl1 FFL
	ffl1.Name = "Wild West Traders"
	ffl1.Address = "123 Austell Rd, Marietta, GA 30166"
	ffl1.Address2 = "123 Austell Rd"
	ffl1.City = "Marietta"
	ffl1.State = "GA"
	ffl1.Zip = "30133"
	ffl1.Country = "USA"
	ffl1.Phone = "800 123 4567"
	ffl1.LicNumber = "456778888"
	return &ffl1
}
