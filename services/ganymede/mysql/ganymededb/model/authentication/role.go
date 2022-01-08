package authentication

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func QueryRolesByActionAndService(ctx context.Context, action, service string) ([]int64, aerror.Error) {
	sqlStr := `SELECT role_id FROM role r
		LEFT JOIN role_action ra ON r.id = ra.role_id
		LEFT JOIN action a ON a.id = ra.action_id
		WHERE a.action = ? AND r.service = ?`

	var roleIds []int64
	err := model.GanymedeDB.SelectContext(ctx, &roleIds, sqlStr, action, service)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "query role_id failed")
	}

	return roleIds, nil
}

func QueryRolesByUserIDAndService(ctx context.Context, userID int64, service string) ([]int64, aerror.Error) {
	sqlStr := `SELECT role_id FROM role r
		LEFT JOIN user_role ur ON  r.id = ur.role_id
		LEFT JOIN user u ON u.id = ur.user_id
		WHERE u.id = ? AND r.service = ?
	`

	var roleIds []int64
	err := model.GanymedeDB.SelectContext(ctx, &roleIds, sqlStr, userID, service)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "query role_id failed")
	}

	return roleIds, nil
}

func QueryRolesByUserIdAndService(ctx context.Context, userID int64, service string) ([]*ganymede.Role, aerror.Error) {
	sqlStr := `SELECT r.id as id, r.name as name , r.service, r.desc FROM role r
		LEFT JOIN user_role ur ON ur.role_id = r.id
		LEFT JOIN user u ON u.id = ur.user_id
		WHERE u.id = ? AND r.service = ?`

	var roles []*ganymede.Role
	err := model.GanymedeDB.SelectContext(ctx, &roles, sqlStr, userID, service)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "query roles failed")
	}

	return roles, nil
}
