package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/product/mysql/charon"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
	"strings"
)

func AttributeMappingInsert(ctx context.Context, mapping *charon.AttributeMapping) (*charon.AttributeMapping, aerror.Error) {
	insertSql := `INSERT INTO attribute_mapping (attr_id, sub_product_id, attr_value) VALUES (:attr_id, :sub_product_id, :attr_value)`
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

func AttributeMappingInsertWithTx(ctx context.Context, tx *sqlx.Tx, mappings []*charon.AttributeMapping) aerror.Error {
	if len(mappings) == 0 {
		return nil
	}
	placeholder := `{{valuePlaceholder}}`
	insertSql := fmt.Sprintf(`INSERT INTO attribute_mapping (attr_id, sub_product_id, attr_value)
				VALUES %s
				ON DUPLICATE KEY UPDATE attr_value=VALUES(attr_value)`, placeholder)

	var values []string
	var params []interface{}
	for _, mapping := range mappings {
		values = append(values, `(?,?,?)`)
		params = append(params, mapping.AttrId, mapping.SubProductId, mapping.AttrValue)
	}

	insertSql = strings.ReplaceAll(insertSql, placeholder, strings.Join(values, ","))

	result, err := tx.ExecContext(ctx, insertSql, params...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	_, err = result.RowsAffected()
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return nil
}

func AttributeMappingQueryById(ctx context.Context, id int64) (*charon.AttributeMapping, aerror.Error) {
	querySql := `SELECT id, attr_id, sub_product_id, attr_value FROM attribute_mapping WHERE id = ?`
	var mappings []*charon.AttributeMapping
	err := charon2.CharonDB.SelectContext(ctx, &mappings, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(mappings) == 0 {
		return nil, nil
	}

	return mappings[0], nil
}

func AttributeMappingQueryBySubProductIDWithTx(ctx context.Context, tx *sqlx.Tx, id int64) ([]*charon.AttributeMapping, aerror.Error) {
	querySql := `SELECT am.id, am.attr_id, am.sub_product_id, am.attr_value
		FROM attribute_mapping am
		LEFT JOIN attribute a ON a.id = am.attr_id
		WHERE a.id IS NOT NULL AND sub_product_id = ?`
	var mappings []*charon.AttributeMapping
	err := charon2.CharonDB.SelectContext(ctx, &mappings, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	return mappings, nil
}

func AttributeMappingDelete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM attribute_mapping WHERE id IN (?)`

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

func AttributeMappingDeleteWithTx(ctx context.Context, tx *sqlx.Tx, ids []int64) aerror.Error {
	delSql := `DELETE FROM attribute_mapping WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del attribute_mapping failed")
	}

	return nil
}

func AttributeMappingDeleteBySubProductIdWithTx(ctx context.Context, tx *sqlx.Tx, subProductIds []int64) aerror.Error {
	if len(subProductIds) == 0 {
		return nil
	}
	delSql := `DELETE FROM attribute_mapping WHERE sub_product_id IN (?)`

	delSql, args, err := sqlx.In(delSql, subProductIds)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del attribute_mapping failed")
	}

	return nil
}

func AttributeMappingDeleteByAttributeIdWithTx(ctx context.Context, tx *sqlx.Tx, attrIds []int64) aerror.Error {
	delSql := `DELETE FROM attribute_mapping WHERE attr_id IN (?)`

	delSql, args, err := sqlx.In(delSql, attrIds)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del attribute_mapping failed")
	}

	return nil
}

func AttributeMappingUpdate(ctx context.Context, product *charon.AttributeMapping) aerror.Error {
	sqlStr := `UPDATE attribute_mapping
		SET attr_id = :attr_id,
		sub_product_id = :sub_product_id,
		attr_value = :attr_value
		WHERE id = :id
`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func AttributeMappingSearch(ctx context.Context, query string, offset, limit int64) (int64, []*charon.AttributeMapping, aerror.Error) {
	formSql := `FROM attribute_mapping `
	whereSql := `WHERE attr_value LIKE ? `
	searchSelect := `SELECT id, attr_id, attr_value, sub_product_id `
	countSelect := `SELECT count(*) `
	query = query + "%"

	countSelectSql := countSelect + formSql + whereSql
	var total int64
	err := charon2.CharonDB.QueryRowx(countSelectSql, query).Scan(&total)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select sub product failed")
	}

	if total == 0 {
		return 0, nil, nil
	}

	searchSelectSql := searchSelect + formSql + whereSql + "limit ?, ?"
	var products []*charon.AttributeMapping
	err = charon2.CharonDB.SelectContext(ctx, &products, searchSelectSql, query, offset, limit)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select sub product failed")
	}

	return total, products, nil
}
