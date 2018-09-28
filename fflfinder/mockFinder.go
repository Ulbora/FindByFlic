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

	var ffl5 FFL
	ffl1.Name = "Bills Arms"
	ffl1.Address = "123 Atlanta Rd, Actworth, GA 30166"
	ffl1.LicNumber = "7894562222"
	rtn = append(rtn, ffl5)

	var ffl6 FFL
	ffl1.Name = "Johns Guns"
	ffl1.Address = "123 Windy Hill Rd, Marietta, GA 30166"
	ffl1.LicNumber = "7894562222"
	rtn = append(rtn, ffl6)
	return &rtn
}
