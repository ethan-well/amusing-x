package session

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"net/http"
	"strings"
)

const CookieName = "user_session"

func SetSession(w http.ResponseWriter, sessionInfo *ganymedeservice.SessionInfo) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    sessionInfo.SessionID,
		Path:     "/",
		MaxAge:   int(sessionInfo.MaxAge),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func GetSessionID(r *http.Request) (string, bool) {
	cook, err := r.Cookie(CookieName)
	if err != nil || cook == nil || len(strings.TrimSpace(cook.Value)) == 0 {
		return "", false
	}

	return strings.TrimSpace(cook.Value), true
}

func DeleteSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}
