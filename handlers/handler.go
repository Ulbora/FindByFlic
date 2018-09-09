package handlers

import (
	"html/template"
	"net/http"
)

//Handler Handler
type Handler struct {
	Templates *template.Template
}

//PageParams PageParams
type PageParams struct {
	Cart  string
	URL   string
	Key   string
	Token string
}

//HandleIndex HandleIndex
func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	var cart = r.URL.Query().Get("cart")
	// 	SecureUrl: 3dcart merchant's Secure URL.
	// PrivateKey: Your application's private key.
	// Token: The 3dcart merchant's token.

	secureURL := r.Header.Get("SecureUrl")
	privateKey := r.Header.Get("PrivateKey")
	token := r.Header.Get("Token")
	var p PageParams
	p.Cart = cart
	p.URL = secureURL
	p.Key = privateKey
	p.Token = token
	h.Templates.ExecuteTemplate(w, "index.html", &p)
}
