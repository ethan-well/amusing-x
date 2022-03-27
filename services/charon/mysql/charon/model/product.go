package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
)

func ProductInsert(ctx context.Context, product *charon.Product) (*charon.Product, aerror.Error) {
	insertSql := `INSERT INTO product (name, description) VALUES (:name, :description)`
	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, product)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	product.ID = id

	return product, nil
}

func ProductQueryById(ctx context.Context, id int64) (*charon.Product, aerror.Error) {
	querySql := `SELECT id, name, description FROM product WHERE id = ?`
	var products []*charon.Product
	err := charon2.CharonDB.SelectContext(ctx, &products, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(products) == 0 {
		return nil, nil
	}

	return products[0], nil
}

func ProductDelete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM product WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del product failed")
	}

	return nil
}

func ProductUpdate(ctx context.Context, product *charon.Product) aerror.Error {
	sqlStr := `UPDATE product SET name = :name, description = :description WHERE id = :id`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func ProductSearch(ctx context.Context, query string, offset, limit int64) ([]*charon.Product, aerror.Error) {
	querySql := `SELECT id, name, description
		FROM product
		WHERE name LIKE ? OR description LIKE ?
		limit ?, ?`

	var products []*charon.Product
	err := charon2.CharonDB.SelectContext(ctx, &products, querySql, "%"+query, "%s"+query, offset, limit)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select product failed")
	}

	return products, nil
}
