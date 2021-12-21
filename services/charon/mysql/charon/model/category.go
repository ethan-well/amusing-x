package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func InsetCategory(ctx context.Context, category *charon.Category) (*charon.Category, aerror.Error) {
	insertSql := `INSERT INTO category (parent_id, name, description) VALUES (:parent_id, :name, :description)`

	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, category)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	category.ID = id

	return category, nil
}
