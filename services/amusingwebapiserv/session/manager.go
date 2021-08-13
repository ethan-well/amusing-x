package session

import (
	"amusingx.fit/amusingx/services/amusingwebapiserv/conf"
	"errors"
	"sync"
)

type Manager struct {
	CookieName  string
	lock        sync.Mutex
	Store       Store
	MaxLifeTime int
}

var GlobalSessionManager *Manager

func InitSessionManager(storeName, cookieName string, maxLiftTTime int) error {
	GlobalSessionManager = &Manager{}
	var err error

	switch storeName {
	case "redis":
		GlobalSessionManager.Store, err = InitRedisStore(conf.Conf.SessionStoreRedisAddr, conf.Conf.SessionStoreRedisPassword,
			conf.Conf.SessionStoreRedisDB, cookieName, maxLiftTTime)
		if err != nil {
			return err
		}
	default:
		return errors.New("session store type is invalid. ")
	}

	GlobalSessionManager.CookieName = cookieName
	GlobalSessionManager.MaxLifeTime = maxLiftTTime

	return nil
}
