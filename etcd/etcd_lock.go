package etcd

import (
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

type XLock struct {
	session *concurrency.Session
	Mutex   *concurrency.Mutex
}

func Lock(key string, timeout time.Duration, ttl int64) (*XLock, aerror.Error) {
	// 生成一个30s超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if Client == nil {
		logger.Errorf("etcd client is nil")
		return nil, aerror.NewErrorf(nil, aerror.Code.SUnexpectedErr, "etcd client is nil")
	}

	// 获取租约
	response, e := Client.Grant(ctx, ttl)
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SUnexpectedErr, "ectd grant failed")
	}

	// 通过租约创建session
	session, e := concurrency.NewSession(Client, concurrency.WithLease(response.ID))
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SUnexpectedErr, "ect concurrency newSession failed")
	}

	mutex := concurrency.NewMutex(session, key)
	e = mutex.Lock(ctx)
	if e != nil {
		// 锁失败，删除当前空租约
		session.Close()
		session.Done()
		return nil, aerror.NewErrorf(e, aerror.Code.SUnexpectedErr, "ectd concurrency newMutex failed")
	}

	return &XLock{
		session: session,
		Mutex:   mutex,
	}, nil
}

func (l *XLock) Unlock(ctx context.Context) aerror.Error {
	defer l.session.Done()
	defer l.session.Close()

	err := l.Mutex.Unlock(ctx)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "etcd unlock failed")
	}

	return nil
}
