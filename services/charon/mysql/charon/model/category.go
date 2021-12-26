package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"database/sql"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
	"strings"
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

func CategoryQuery(ctx context.Context, id int64, name, desc string, offSet, limit int64) ([]*charon.Category, int64, aerror.Error) {
	wherePlaceholder := `{{whereCondition}}`
	sqlStr := fmt.Sprintf(`SELECT id, name, description FROM category  %s limit ?, ?`, wherePlaceholder)
	countStr := fmt.Sprintf(`SELECT count(*) FROM category  %s limit ?, ?`, wherePlaceholder)

	var whereConditions []string
	var params []interface{}
	if id > 0 {
		whereConditions = append(whereConditions, `id = ?`)
		params = append(params, id)
	}

	if len(name) != 0 {
		whereConditions = append(whereConditions, "name = ?")
		params = append(params, name)
	}

	if len(desc) != 0 {
		whereConditions = append(whereConditions, "description = ?")
		params = append(params, desc)
	}

	params = append(params, offSet, limit)

	if len(whereConditions) != 0 {
		whereSQl := "WHERE " + strings.Join(whereConditions, " AND ")
		sqlStr = strings.Replace(sqlStr, wherePlaceholder, whereSQl, -1)
		countStr = strings.Replace(countStr, wherePlaceholder, whereSQl, -1)
	} else {
		sqlStr = strings.Replace(sqlStr, wherePlaceholder, "", -1)
		countStr = strings.Replace(countStr, wherePlaceholder, "", -1)
	}

	tx, err := charon2.CharonDB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "begin transaction failed")
	}

	var count int64
	err = tx.QueryRowx(countStr, params...).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		count = 0
	case err != nil:
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query count error")
	}

	var categories []*charon.Category
	err = tx.SelectContext(ctx, &categories, sqlStr, params...)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select total category list failed")
	}

	return categories, count, nil
}

func Delete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM category WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del category failed")
	}

	return nil
}

func QueryCategoryByID(ctx context.Context, id int64) (*charon.Category, aerror.Error) {
	sqlStr := `SELECT id, name, description FROM category WHERE id = ?`

	var category charon.Category
	err := charon2.CharonDB.GetContext(ctx, &category, sqlStr, id)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select failed")
	}

	return &category, nil
}
