package session

import (
	"net/http"
	"strings"
)

func GetSessionID(r *http.Request) (string, bool) {
	cook, err := r.Cookie(CookieName)
	if err != nil || cook == nil || len(strings.TrimSpace(cook.Value)) == 0 {
		return "", false
	}

	return strings.TrimSpace(cook.Value), true
}
