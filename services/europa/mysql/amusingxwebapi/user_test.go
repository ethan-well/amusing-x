package amusingxwebapi

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/europa/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestQueryUserByIdContext(t *testing.T) {
	conf.Mock()
	Mock()

	user, err := QueryUserByIdContext(context.TODO(), 1)
	if err != nil {
		t.Fatalf("some error occurred: %s", err)
	}

	if user == nil {
		t.Logf("no user with id = %d", 1)
	} else {
		t.Logf("succced, user id: %d", user.ID)
	}
}

func TestQueryUserByNicknameOrPhone(t *testing.T) {
	conf.Mock()
	Mock()

	u, err := QueryUserByNicknameOrPhone(context.Background(), "wei.wei", "18710565589")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	t.Logf("u: %v", u)
}

func TestInsert(t *testing.T) {
	conf.Mock()
	Mock()

	u := &ganymede.UserComplex{
		Nickname:       "wei.wei3",
		Phone:          "18710565582",
		PasswordDigest: "password digest test",
		Salt:           "salt test",
	}

	u, err := Insert(context.TODO(), u)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("user: %s", logger.ToJson(u))
}

func TestQueryUserByPhone(t *testing.T) {
	conf.Mock()
	Mock()

	user, err := QueryUserByPhone(context.TODO(), "86-18710565585")
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("user: %s", logger.ToJson(user))
}
