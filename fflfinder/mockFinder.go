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
	ffl1.LicenseName = "Wild West Traders"
	ffl1.PremiseStreet = "123 Austell Rd"
	ffl1.PremiseCity = "Marietta"
	ffl1.PremiseState = "GA"
	ffl1.PremiseZipCode = "30133"
	ffl1.VoicePhone = "800 123 4567"
	rtn = append(rtn, ffl1)

	var ffl2 FFL
	ffl2.LicenseName = "Wild West Traders 2"
	ffl2.PremiseStreet = "123 Austell Rd"
	ffl2.PremiseCity = "Marietta"
	ffl2.PremiseState = "GA"
	ffl2.PremiseZipCode = "30133"
	ffl2.VoicePhone = "800 123 4567"
	rtn = append(rtn, ffl2)

	var ffl3 FFL
	ffl3.LicenseName = "Wild West Traders 3"
	ffl3.PremiseStreet = "123 Austell Rd"
	ffl3.PremiseCity = "Marietta"
	ffl3.PremiseState = "GA"
	ffl3.PremiseZipCode = "30133"
	ffl3.VoicePhone = "800 123 4567"
	rtn = append(rtn, ffl3)

	var ffl4 FFL
	ffl4.LicenseName = "Wild West Traders 4"
	ffl4.PremiseStreet = "123 Austell Rd"
	ffl4.PremiseCity = "Marietta"
	ffl4.PremiseState = "GA"
	ffl4.PremiseZipCode = "30133"
	ffl4.VoicePhone = "800 123 4567"
	rtn = append(rtn, ffl4)

	var ffl5 FFL
	ffl5.LicenseName = "Wild West Traders 5"
	ffl5.PremiseStreet = "123 Austell Rd"
	ffl5.PremiseCity = "Marietta"
	ffl5.PremiseState = "GA"
	ffl5.PremiseZipCode = "30133"
	ffl5.VoicePhone = "800 123 4567"
	rtn = append(rtn, ffl5)

	var ffl6 FFL
	ffl6.LicenseName = "Wild West Traders 5"
	ffl6.PremiseStreet = "123 Austell Rd"
	ffl6.PremiseCity = "Marietta"
	ffl6.PremiseState = "GA"
	ffl6.PremiseZipCode = "30133"
	ffl6.VoicePhone = "800 123 4567"
	rtn = append(rtn, ffl6)
	return &rtn
}

//GetFFL GetFFL
func (f *MockFinder) GetFFL(id int64) *FFL {
	var ffl1 FFL
	ffl1.LicenseName = "Wild West Traders"
	ffl1.PremiseStreet = "123 Austell Rd"
	ffl1.PremiseCity = "Marietta"
	ffl1.PremiseState = "GA"
	ffl1.PremiseZipCode = "30133"
	ffl1.VoicePhone = "800 123 4567"
	return &ffl1
}
