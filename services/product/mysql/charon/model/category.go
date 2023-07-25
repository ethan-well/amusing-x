package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/product/mysql/charon"
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

func CategoryQuery(ctx context.Context, req *proto.CategoryListRequest, filter *charonservice.SearchFilter) ([]*charon.Category, int64, aerror.Error) {
	wherePlaceholder := `{{whereCondition}}`
	sqlStr := fmt.Sprintf(`SELECT id, name, description FROM category %s limit ?, ?`, wherePlaceholder)
	countStr := fmt.Sprintf(`SELECT count(*) FROM category  %s`, wherePlaceholder)

	var whereConditions []string
	var params []interface{}
	if len(filter.Id) > 0 {
		whereConditions = append(whereConditions, `id in (?)`)
		params = append(params, filter.Id)
	}

	if len(filter.Name) > 0 {
		whereConditions = append(whereConditions, "name in in (?)")
		params = append(params, filter.Name)
	}

	if len(filter.Desc) != 0 {
		whereConditions = append(whereConditions, "description in (?)")
		params = append(params, filter.Desc)
	}

	params = append(params, req.Offset, req.Limit)

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
	defer tx.Rollback()

	var count int64
	countStr, args, err := sqlx.In(countStr, params[:len(params)-2]...)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err = tx.QueryRowx(countStr, args...).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		count = 0
	case err != nil:
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query count error")
	}

	var categories []*charon.Category
	sqlStr, args, err = sqlx.In(sqlStr, params...)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err = tx.SelectContext(ctx, &categories, sqlStr, args...)
	if err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select total category list failed")
	}

	if err := tx.Commit(); err != nil {
		return nil, 0, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql commit error")
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

func UpdateCategory(ctx context.Context, category *charon.Category) aerror.Error {
	sqlStr := `UPDATE category SET name = :name, description = :description WHERE id = :id`

	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, category)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func QueryCategoryByName(ctx context.Context, name string) (*charon.Category, aerror.Error) {
	sqlStr := `SELECT id, name, description FROM category WHERE name = ?`

	var category charon.Category
	err := charon2.CharonDB.GetContext(ctx, &category, sqlStr, name)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select failed")
	}

	return &category, nil
}

func SearchCategoryByProductIds(ctx context.Context, productIds []int64) ([]*charon.CategoryWide, aerror.Error) {
	querySql := `SELECT c.id as id, c.parent_id as parent_id, p.id as product_id, c.name as name, c.description as description
		FROM category c
		LEFT JOIN category_product_mapping map on c.id = map.category_id
		LEFT JOIN product p on p.id = map.product_id
		WHERE p.id in (?) AND c.id > 0`

	querySql, args, err := sqlx.In(querySql, productIds)
	if err != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	var categories []*charon.CategoryWide
	err = charon2.CharonDB.SelectContext(ctx, &categories, querySql, args...)
	if err != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "search category by product_id failed")
	}

	return categories, nil
}
