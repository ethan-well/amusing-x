package authentication

import (
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/mysql/ganymededb/model"
	"context"
	"testing"
)

func TestHavePermission(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	conf.Mock()
	model.Mock()

	var tester = map[bool]map[int64]string{
		true:  {21: "all-action"},
		false: {21: "no-action"},
	}

	for e, r := range tester {
		for u, a := range r {
			ok, _ := Permission(context.Background(), u, a, "pangu")
			if ok != e {
				t.Logf("expected: %t, get: %t", e, ok)
			}
		}
	}
}
