package authentication

import (
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func QueryRolesByAction(ctx context.Context, action string) ([]int64, aerror.Error) {
	sqlStr := `SELECT role_id FROM role r
		LEFT JOIN role_action ra ON r.id = ra.role_id
		LEFT JOIN action a ON a.id = ra.action_id
		WHERE a.action = ?`

	var roleIds []int64
	err := model.GanymedeDB.SelectContext(ctx, &roleIds, sqlStr, action)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "query role_id failed")
	}

	return roleIds, nil
}

func QueryRolesByUserID(ctx context.Context, userID int64) ([]int64, aerror.Error) {
	sqlStr := `SELECT role_id FROM role r
		LEFT JOIN user_role ur ON  r.id = ur.role_id
		LEFT JOIN user u ON u.id = ur.user_id
		WHERE u.id = ?
	`

	var roleIds []int64
	err := model.GanymedeDB.SelectContext(ctx, &roleIds, sqlStr, userID)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "query role_id failed")
	}

	return roleIds, nil
}
