package session

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
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

var MaxAge = 7 * 24 * 60 * 60 //second

func InitSessionManager(storeName string, maxLiftTTime int) error {
	GlobalSessionManager = &Manager{}
	var err error

	switch storeName {
	case "redis":
		redis := conf.Conf.SessionStore.Redis
		GlobalSessionManager.Store, err = InitRedisStore(redis.Addr, redis.Password, redis.DBNo, conf.Conf.SessionStore.Prefix, maxLiftTTime)
		if err != nil {
			return err
		}
	default:
		return errors.New("session store type is invalid. ")
	}

	GlobalSessionManager.CookieName = conf.Conf.SessionStore.Prefix
	GlobalSessionManager.MaxLifeTime = maxLiftTTime

	return nil
}
