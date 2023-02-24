package rawebclients

import (
	tcp "github.com/ilinovalex86/ratcpserver"
	"net/http"
)

// Valid - Проверяет cookie web клиента + возвращает индекс tcp клиента
func Valid(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("SessionId")
	if err == nil {
		webId := cookie.Value
		if tcpId, ok := Clients.Load(webId); ok {
			if tcp.Clients.Exist(tcpId) {
				return tcpId, true
			}
		}
	}
	return "", false
}
