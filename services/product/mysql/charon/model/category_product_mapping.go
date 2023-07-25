package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/product/mysql/charon"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
)

func CategoryProductMappingInsert(ctx context.Context, mapping *charon.CategoryProductMapping) (*charon.CategoryProductMapping, aerror.Error) {
	insertSql := `INSERT INTO category_product_mapping (category_id, product_id) VALUES (:category_id, :product_id)`
	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, mapping)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	mapping.ID = id

	return mapping, nil
}

func CategoryProductMappingInsertWithTx(ctx context.Context, tx *sqlx.Tx, mapping *charon.CategoryProductMapping) (*charon.CategoryProductMapping, aerror.Error) {
	insertSql := `INSERT INTO category_product_mapping (category_id, product_id) VALUES (:category_id, :product_id)`
	result, err := tx.NamedExecContext(ctx, insertSql, mapping)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	mapping.ID = id

	return mapping, nil
}

func CategoryProductMappingQueryById(ctx context.Context, id int64) (*charon.CategoryProductMapping, aerror.Error) {
	querySql := `SELECT id, category_id, product_id, create_time, update_time FROM category_product_mapping WHERE id = ?`
	var mapping []*charon.CategoryProductMapping
	err := charon2.CharonDB.SelectContext(ctx, &mapping, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(mapping) == 0 {
		return nil, nil
	}

	return mapping[0], nil
}

func CategoryProductMappingDelete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM category_product_mapping WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "delete category failed")
	}

	return nil
}

func CategoryProductMappingDeleteWithTX(ctx context.Context, tx *sql.Tx, ids []int64) aerror.Error {
	delSql := `DELETE FROM category_product_mapping WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "delete category failed")
	}

	return nil
}

func CategoryProductMappingDeleteByProductIdWithTX(ctx context.Context, tx *sqlx.Tx, productIds []int64) aerror.Error {
	delSql := `DELETE FROM category_product_mapping WHERE product_id in (?)`

	delSql, args, err := sqlx.In(delSql, productIds)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "delete category failed")
	}

	return nil
}

func CategoryProductMappingUpdate(ctx context.Context, mapping *charon.CategoryProductMapping) aerror.Error {
	sqlStr := `UPDATE category_product_mapping
		SET category_id = :category_id,
		product_id = :product_id
		WHERE id = :id
`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, mapping)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func CategoryProductMappingUpdateByProductIdWithTx(ctx context.Context, tx *sqlx.Tx, mapping *charon.CategoryProductMapping) aerror.Error {
	sqlStr := `UPDATE category_product_mapping
		SET category_id = :category_id,
		product_id = :product_id
		WHERE product_id = :product_id
`
	_, err := tx.NamedExecContext(ctx, sqlStr, mapping)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}
