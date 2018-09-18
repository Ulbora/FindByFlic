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

package dbdelegate

const (
	dcartTest        = "select count(*) from dcart_user "
	dcartGetByStore  = "select * from dcart_user where store_url = ? "
	dcartInsertStore = "insert into dcart_user (store_url, public_key, token, action, entered, enabled) values(?, ?, ?, ?, ?, ?)"
	dcartUpdateStore = "update dcart_user set public_key = ?, token = ?, action = ?, modified = ?, enabled = true where id = ? "
	dcartRemoveStore = "update dcart_user set action = ?, modified = ?, enabled = false where id = ? "
)
