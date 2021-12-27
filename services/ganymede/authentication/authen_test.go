package authentication

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
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
			ok, _ := HavePermission(context.Background(), u, a)
			if ok != e {
				t.Logf("expected: %t, get: %t", e, ok)
			}
		}
	}
}
