package session

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"errors"
	"sync"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	store       Store
	maxLifeTime int64
}

func NewManager(storeName, cookieName string, maxLiftTTime int64) (*Manager, error) {
	manager := &Manager{}
	var err error

	switch storeName {
	case "redis":
		manager.store, err = InitRedisStore(conf.Conf.SessionStoreRedisAddr, conf.Conf.SessionStoreRedisPassword,
			conf.Conf.SessionStoreRedisDB, cookieName, maxLiftTTime)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("session store type is invalid. ")
	}

	manager.cookieName = cookieName
	manager.maxLifeTime = maxLiftTTime

	return manager, nil
}
