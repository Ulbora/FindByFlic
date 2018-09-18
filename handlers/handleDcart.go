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

package handlers

import (
	"log"
	"net/http"
)

//HandleDcartIndex HandleDcartIndex
func (h *Handler) HandleDcartIndex(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	var cart = r.URL.Query().Get("cart")
	var carturl = r.URL.Query().Get("carturl")
	session := h.getSession(w, r)
	session.Values["cart"] = cart
	session.Values["carturl"] = carturl
	serr := session.Save(r, w)
	log.Println(serr)
	// 	SecureUrl: 3dcart merchant's Secure URL.
	// PrivateKey: Your application's private key.
	// Token: The 3dcart merchant's token.
	//<iframe src="https://localhost:8070?cart=3dcar&carturl=[store_url]"></iframe>
	//<iframe src="https://localhost:8070?cart=3dcar&carturl=[store_url]"></iframe>

	// secureURL := r.Header.Get("SecureUrl")
	// privateKey := r.Header.Get("PrivateKey")
	// token := r.Header.Get("Token")
	var p PageParams
	p.Cart = cart
	p.CartURL = carturl
	// p.URL = secureURL
	// p.Key = privateKey
	// p.Token = token
	h.Templates.ExecuteTemplate(w, "dcartIndex.html", &p)
}
