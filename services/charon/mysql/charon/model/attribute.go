package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
)

func AttributeInsert(ctx context.Context, attribute *charon.Attribute) (*charon.Attribute, aerror.Error) {
	insertSql := `INSERT INTO attribute (name, description) VALUES (:name, :description)`
	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, attribute)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	attribute.ID = id

	return attribute, nil
}

func AttributeQueryById(ctx context.Context, id int64) (*charon.Attribute, aerror.Error) {
	querySql := `SELECT name, description FROM attribute WHERE id = ?`
	var products []*charon.Attribute
	err := charon2.CharonDB.SelectContext(ctx, &products, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(products) == 0 {
		return nil, nil
	}

	return products[0], nil
}

func AttributeDelete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM attribute WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del attribute failed")
	}

	return nil
}

func AttributeUpdate(ctx context.Context, product *charon.Attribute) aerror.Error {
	sqlStr := `UPDATE attribute
		SET name = :name,
		description = :description
		WHERE id = :id
`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func AttributeSearch(ctx context.Context, query string, offset, limit int64) (int64, []*charon.Attribute, aerror.Error) {
	formSql := `FROM attribute `
	whereSql := `WHERE name LIKE ? OR description LIKE ? `
	searchSelect := `SELECT id, name, description `
	countSelect := `SELECT count(*) `
	query = query + "%"

	countSelectSql := countSelect + formSql + whereSql
	var total int64
	err := charon2.CharonDB.QueryRowx(countSelectSql, query, query).Scan(&total)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select attribute failed")
	}

	if total == 0 {
		return 0, nil, nil
	}

	searchSelectSql := searchSelect + formSql + whereSql + "limit ?, ?"
	var products []*charon.Attribute
	err = charon2.CharonDB.SelectContext(ctx, &products, searchSelectSql, query, query, offset, limit)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select attribute failed")
	}

	return total, products, nil
}
