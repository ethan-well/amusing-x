package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"database/sql"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/jmoiron/sqlx"
	"strings"
)

func SubProductInsert(ctx context.Context, subProduct *charon.SubProduct) (*charon.SubProduct, aerror.Error) {
	insertSql := `INSERT INTO sub_product (product_id, name, description, currency, price, stock) VALUES (:product_id, :name, :description, :currency, :price, :stock)`
	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, subProduct)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	subProduct.ID = id

	return subProduct, nil
}

func SubProductQueryById(ctx context.Context, id int64) (*charon.SubProduct, aerror.Error) {
	querySql := `SELECT id, product_id, name, description, currency, price, stock FROM sub_product WHERE id = ?`
	var products []*charon.SubProduct
	err := charon2.CharonDB.SelectContext(ctx, &products, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(products) == 0 {
		return nil, nil
	}

	return products[0], nil
}

func SubProductQueryByIdWithTx(ctx context.Context, id int64, tx ...*sqlx.Tx) (*charon.SubProduct, aerror.Error) {
	querySql := `SELECT id, product_id, name, description, currency, price, stock FROM sub_product WHERE id = ?`
	var products []*charon.SubProduct
	var err error
	if tx == nil {
		err = charon2.CharonDB.SelectContext(ctx, &products, querySql, id)
	} else {
		err = tx[0].SelectContext(ctx, &products, querySql, id)
	}
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(products) == 0 {
		return nil, nil
	}

	return products[0], nil
}

func SubProductDelete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM sub_product WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del sub product failed")
	}

	return nil
}

func SubProductDeleteWithTx(ctx context.Context, tx *sqlx.Tx, ids []int64) aerror.Error {
	delSql := `DELETE FROM sub_product WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del sub_product failed")
	}

	return nil
}

func SubProductUpdate(ctx context.Context, product *charon.SubProduct) aerror.Error {
	sqlStr := `UPDATE sub_product
		SET name = :name,
		description = :description,
		product_id = :product_id,
		currency = :currency,
		price = :price,
		stock = :stock
		
		WHERE id = :id
`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func SubProductUpdateWithTx(ctx context.Context, tx *sqlx.Tx, product *charon.SubProduct) aerror.Error {
	sqlStr := `UPDATE sub_product
		SET name = :name,
		description = :description,
		product_id = :product_id,
		currency = :currency,
		price = :price,
		stock = :stock
		WHERE id = :id
`
	_, err := tx.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func SubProductSearch(ctx context.Context, query string, offset, limit int64) (int64, []*charon.SubProduct, aerror.Error) {
	formSql := `FROM sub_product `
	whereSql := `WHERE name LIKE ? OR description LIKE ? `
	searchSelect := `SELECT id, name, description, product_id, currency, price, stock `
	countSelect := `SELECT count(*) `
	query = query + "%"

	countSelectSql := countSelect + formSql + whereSql
	var total int64
	err := charon2.CharonDB.QueryRowx(countSelectSql, query, query).Scan(&total)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select sub product failed")
	}

	if total == 0 {
		return 0, nil, nil
	}

	searchSelectSql := searchSelect + formSql + whereSql + "limit ?, ?"
	var products []*charon.SubProduct
	err = charon2.CharonDB.SelectContext(ctx, &products, searchSelectSql, query, query, offset, limit)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select sub product failed")
	}

	return total, products, nil
}

func SubProductSearchV2(ctx context.Context, in *proto.SubProductListRequest, filter *charonservice.SearchFilter) (int64, []*charon.SubProduct, aerror.Error) {
	wherePlaceholder := `{{whereCondition}}`
	sqlStr := fmt.Sprintf(`SELECT id, name, description, product_id, currency, price, stock FROM sub_product %s limit ?, ?`, wherePlaceholder)
	countStr := fmt.Sprintf(`SELECT count(*) FROM sub_product %s`, wherePlaceholder)

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

	params = append(params, in.Offset, in.Limit)

	if len(whereConditions) != 0 {
		whereSQl := "WHERE " + strings.Join(whereConditions, " AND ")
		sqlStr = strings.Replace(sqlStr, wherePlaceholder, whereSQl, -1)
		countStr = strings.Replace(countStr, wherePlaceholder, whereSQl, -1)
	} else {
		sqlStr = strings.Replace(sqlStr, wherePlaceholder, "", -1)
		countStr = strings.Replace(countStr, wherePlaceholder, "", -1)
	}

	logger.Errorf("sqlStr: %s", sqlStr)
	logger.Errorf("countStr: %s", countStr)
	logger.Errorf("params: %v", params)

	tx, err := charon2.CharonDB.BeginTxx(ctx, nil)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "begin transaction failed")
	}
	defer tx.Rollback()

	var count int64
	countStr, args, err := sqlx.In(countStr, params[:len(params)-2]...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err = tx.QueryRowx(countStr, args...).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		count = 0
	case err != nil:
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query count error")
	}

	var subProducts []*charon.SubProduct
	sqlStr, args, err = sqlx.In(sqlStr, params...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err = tx.SelectContext(ctx, &subProducts, sqlStr, args...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select total category list failed")
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql commit error")
	}

	return count, subProducts, nil
}
