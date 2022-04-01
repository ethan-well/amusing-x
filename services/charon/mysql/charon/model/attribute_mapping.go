package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
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
