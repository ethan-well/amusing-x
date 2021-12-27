package authentication

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"testing"
)

func TestQueryRolesByAction(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	conf.Mock()
	model.Mock()

	roleIds, err := QueryRolesByAction(context.Background(), "all-action")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("roleIds: %v", roleIds)
}

func TestQueryRolesByUserID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	conf.Mock()
	model.Mock()

	roleIds, err := QueryRolesByUserID(context.Background(), 21)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("roleIds: %v", roleIds)
}
