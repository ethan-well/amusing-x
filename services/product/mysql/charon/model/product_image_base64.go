package model

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charon2 "amusingx.fit/amusingx/services/product/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/jmoiron/sqlx"
)

func ProductImageBase64Insert(ctx context.Context, image *charon.ProductImageBase64) (*charon.ProductImageBase64, aerror.Error) {
	insertSql := `INSERT INTO product_image_base64 (base64) VALUES (:base64)`
	result, err := charon2.CharonDB.NamedExecContext(ctx, insertSql, image)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	image.Id = id

	return image, nil
}

func ProductImageBase64InsertWithTx(ctx context.Context, tx *sqlx.Tx, image *charon.ProductImageBase64) (*charon.ProductImageBase64, aerror.Error) {
	insertSql := `INSERT INTO product_image_base64 (base64) VALUES (:base64)`
	result, err := tx.NamedExecContext(ctx, insertSql, image)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "get last insert id failed")
	}

	image.Id = id

	return image, nil
}

func ProductImageBase64QueryById(ctx context.Context, id int64) (*charon.ProductImageBase64, aerror.Error) {
	var image charon.ProductImageBase64
	insertSql := `SELECT id, base64 FROM product_image_base64 WHERE id = ?`

	err := charon2.CharonDB.GetContext(ctx, &image, insertSql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return &image, nil
}

func ProductImageBase64Delete(ctx context.Context, ids []int64) aerror.Error {
	insertSql := `DELETE FROM product_image_base64 where id in (?)`
	insertSql, args, e := sqlx.In(insertSql, ids)
	if e != nil {
		return aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err := charon2.CharonDB.ExecContext(ctx, insertSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return nil
}
