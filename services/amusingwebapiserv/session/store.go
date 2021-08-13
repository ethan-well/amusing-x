package session

import "context"

type Store interface {
	SessionInit(ctx context.Context, sid string) (Session, error)
	SessionRead(ctx context.Context, sid string) (Session, error)
	SessionDestroy(ctx context.Context, sid string) error
	SessionGC(ctx context.Context, maxLifeTime int64)
	SessionWrite(ctx context.Context, sid string, value interface{}) error
}
