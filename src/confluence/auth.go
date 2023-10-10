package confluence

import "net/http"

func (a *API) Authenticate(req *http.Request) {
	if a.username != "" && a.token != "" {
		req.SetBasicAuth(a.username, a.token)
	} else if a.token != "" {
		req.Header.Set("Authorization", "Bearer "+a.token)
	}
}
