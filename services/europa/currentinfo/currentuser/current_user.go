package currentuser

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"sync"
)

type CurrentUser ganymedeservice.UserInfo

var (
	currentUser *CurrentUser
	mux         sync.RWMutex
)

func GetCurrentUser() *CurrentUser {
	mux.RLock()
	mux.RUnlock()
	return currentUser
}

func SetCurrentUser(user *CurrentUser) {
	mux.Lock()
	defer mux.Unlock()
	currentUser = user
}

func ClearCurrentUser() {
	mux.Lock()
	defer mux.Unlock()

	currentUser = nil
}
