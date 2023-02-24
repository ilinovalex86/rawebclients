package rawebclients

import "net/http"

// ClientClear - Обнуляет cookie у web клиента и перенаправляет на login
func ClientClear(w http.ResponseWriter, r *http.Request) {
	cookieR, err := r.Cookie("SessionId")
	if err == nil {
		webId := cookieR.Value
		if Clients.exist(webId) {
			Clients.delete(webId)
		}
	}
	cookie := http.Cookie{Name: "SessionId", Value: "-"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}
