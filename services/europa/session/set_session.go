package session

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"net/http"
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
