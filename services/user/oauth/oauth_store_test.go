package oauth

import (
	"amusingx.fit/amusingx/services/user/conf"
	"context"
	"testing"
)

func TestStateCode(t *testing.T) {
	conf.Mock()
	store := conf.Conf.OAuth.Store
	redis := store.Redis
	InitRedisStore(redis.Addr, redis.Password, redis.DBNo, store.Prefix, store.MaxLifeTime)

	ctx := context.Background()
	stateCode, err := StoreIns.RandomStateCode(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("stateCode", stateCode)

	err = StoreIns.ValidSate(ctx, stateCode)
	if err != nil {
		t.Fatal(err)
	}

	err = StoreIns.ValidSate(ctx, "273678")
	if err == nil {
		t.Fatal(err)
	}
	t.Logf("succeed")
}
