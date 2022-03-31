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

func ProductUpdate(ctx context.Context, product *charon.Product) aerror.Error {
	sqlStr := `UPDATE product SET name = :name, description = :description WHERE id = :id`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
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
