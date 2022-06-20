package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
)

func ProductStockInsert(ctx context.Context, productStock *charon.ProductStock, tx ...*sqlx.Tx) (*charon.ProductStock, aerror.Error) {
	insertSql := `INSERT INTO product_stock (sub_product_id, real_inventory, available_inventory)
		VALUES (:sub_product_id, :real_inventory, :available_inventory)`

	var (
		result sql.Result
		err    error
	)
	if tx != nil {
		result, err = tx[0].NamedExecContext(ctx, insertSql, productStock)
	} else {
		result, err = charon2.CharonDB.NamedExecContext(ctx, insertSql, productStock)
	}

	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	productStock.ID = id

	return productStock, nil
}

func QueryProductStockBySubProductIds(ctx context.Context, ids []int64, tx ...*sqlx.Tx) ([]*charon.ProductStock, aerror.Error) {
	querySql := `SELECT id, sub_product_id, real_inventory, available_inventory
		FROM product_stock where sub_product_id in (?)`

	querySql, args, err := sqlx.In(querySql, ids)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	var productStocks []*charon.ProductStock
	if tx != nil {
		err = tx[0].SelectContext(ctx, &productStocks, querySql, args...)
	} else {
		err = charon2.CharonDB.SelectContext(ctx, &productStocks, querySql, args...)
	}
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select failed")
	}

	return productStocks, nil
}

func ProductStockDeleteBySubProductIds(ctx context.Context, ids []int64, tx ...*sqlx.Tx) aerror.Error {
	delSql := `DELETE FROM product_stock WHERE sub_product_id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	if tx != nil {
		_, err = tx[0].ExecContext(ctx, delSql, args...)
	} else {
		charon2.CharonDB.ExecContext(ctx, delSql, args...)
	}

	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del sub_product failed")
	}

	return nil
}

func ProductStockUpdateWithTx(ctx context.Context, productStock *charon.ProductStock, tx ...*sqlx.Tx) aerror.Error {
	sqlStr := `UPDATE product_stock
		SET real_inventory = :real_inventory,
			available_inventory = :available_inventory
		WHERE sub_product_id = :sub_product_id
	`
	var err error
	if tx != nil {
		_, err = tx[0].NamedExecContext(ctx, sqlStr, productStock)
	} else {
		_, err = charon2.CharonDB.NamedExecContext(ctx, sqlStr, productStock)
	}

	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}
