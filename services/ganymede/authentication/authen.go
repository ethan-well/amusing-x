package authentication

import (
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model/authentication"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/set/intset"
)

func Permission(ctx context.Context, userID int64, action, service string) (bool, aerror.Error) {
	actionRoles, err := authentication.QueryRolesByActionAndService(ctx, action, service)
	if err != nil {
		return false, err
	}

	if len(actionRoles) == 0 {
		return false, nil
	}

	userRoles, err := authentication.QueryRolesByUserIDAndService(ctx, userID, service)
	if err != nil {
		return false, err
	}

	if len(userRoles) == 0 {
		return false, nil
	}

	set1 := intset.NewInt64Set(actionRoles...)
	set2 := intset.NewInt64Set(userRoles...)

	if len(set1.Intersection(set2)) != 0 {
		return true, nil
	}

	return false, nil
}
