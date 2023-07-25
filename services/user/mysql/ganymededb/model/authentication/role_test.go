package authentication

import (
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/mysql/ganymededb/model"
	"context"
	"testing"
)

func TestQueryRolesByAction(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	conf.Mock()
	model.Mock()

	roleIds, err := QueryRolesByActionAndService(context.Background(), "all-action", "pangu")
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

	roleIds, err := QueryRolesByUserIDAndService(context.Background(), 21, "pangu")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("roleIds: %v", roleIds)
}

func TestQueryRolesByUserIdAndService(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	conf.Mock()
	model.Mock()

	roles, err := QueryRolesByUserIdAndService(context.Background(), 21, "pangu")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("roles: %#v", roles[0])
}
