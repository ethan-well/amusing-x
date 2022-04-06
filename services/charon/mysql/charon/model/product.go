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
	"github.com/jmoiron/sqlx"
	"strings"
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

func ProductInsertWithTx(ctx context.Context, tx *sqlx.Tx, product *charon.Product) (*charon.Product, aerror.Error) {
	insertSql := `INSERT INTO product (name, description) VALUES (:name, :description)`
	result, err := tx.NamedExecContext(ctx, insertSql, product)
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

func ProductWideInfoById(ctx context.Context, id int64) (*charon.ProductWide, aerror.Error) {
	querySql := `SELECT p.id as id,
    	IFNULL(c.id, 0) as category_id,
    	p.name as name,
       	p.description as description,
		p.create_time as create_time,
		p.update_time as update_time
	FROM product p
	LEFT JOIN category_product_mapping map ON map.product_id = p.id
	LEFT JOIN category c ON map.category_id = c.id
	WHERE p.id = ?;`

	var products []*charon.ProductWide
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

func ProductDeleteWithTx(ctx context.Context, tx *sqlx.Tx, ids []int64) aerror.Error {
	delSql := `DELETE FROM product WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
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

func ProductUpdateWithTx(ctx context.Context, tx *sqlx.Tx, product *charon.Product) aerror.Error {
	sqlStr := `UPDATE product SET name = :name, description = :description WHERE id = :id`
	_, err := tx.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func ProductSearchQuery(ctx context.Context, req *proto.ProductListRequest, filter *charonservice.SearchFilter) (int64, []*charon.Product, aerror.Error) {
	wherePlaceholder := `{{whereCondition}}`
	sqlStr := fmt.Sprintf(`SELECT id, name, description FROM product %s limit ?, ?`, wherePlaceholder)
	countStr := fmt.Sprintf(`SELECT count(*) FROM product  %s`, wherePlaceholder)

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

	var categories []*charon.Product
	sqlStr, args, err = sqlx.In(sqlStr, params...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err = tx.SelectContext(ctx, &categories, sqlStr, args...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select total category list failed")
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql commit error")
	}

	return count, categories, nil
}

func ProductSearch(ctx context.Context, query string, offset, limit int64) (int64, []*charon.Product, aerror.Error) {
	formSql := `FROM product `
	whereSql := `WHERE name LIKE ? OR description LIKE ? `
	searchSelect := `SELECT id, name, description `
	countSelect := `SELECT count(*) `
	query = query + "%"

	countSelectSql := countSelect + formSql + whereSql
	var total int64
	err := charon2.CharonDB.QueryRowx(countSelectSql, query, query).Scan(&total)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select product failed")
	}

	if total == 0 {
		return 0, nil, nil
	}

	searchSelectSql := searchSelect + formSql + whereSql + "limit ?, ?"
	var products []*charon.Product
	err = charon2.CharonDB.SelectContext(ctx, &products, searchSelectSql, query, query, offset, limit)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select product failed")
	}

	return total, products, nil
}
