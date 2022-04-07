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

func ProductImageInsert(ctx context.Context, image *charon.ProductImage) (*charon.ProductImage, aerror.Error) {
	insertSql := `INSERT INTO product_image (product_id, product_level, uploader_type, url) VALUES (:product_id, :product_level, :uploader_type, :url)`
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

func ProductImageQueryById(ctx context.Context, id int64) (*charon.ProductImage, aerror.Error) {
	querySql := `SELECT id, product_id, product_level, uploader_type, url FROM product_image WHERE id = ?`
	var images []*charon.ProductImage
	err := charon2.CharonDB.SelectContext(ctx, &images, querySql, id)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(images) == 0 {
		return nil, nil
	}

	return images[0], nil
}

func ProductImageQueryByIdWithTx(ctx context.Context, id int64, tx ...*sqlx.Tx) (*charon.ProductImage, aerror.Error) {
	querySql := `SELECT id, product_id, product_level, uploader_type, url FROM product_image WHERE id = ?`
	var images []*charon.ProductImage
	var err error
	if tx == nil {
		err = charon2.CharonDB.SelectContext(ctx, &images, querySql, id)
	} else {
		err = tx[0].SelectContext(ctx, &images, querySql, id)
	}
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}
	if len(images) == 0 {
		return nil, nil
	}

	return images[0], nil
}

func ProductImageDelete(ctx context.Context, ids []int64) aerror.Error {
	delSql := `DELETE FROM product_image WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = charon2.CharonDB.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del product_image failed")
	}

	return nil
}

func ProductImageDeleteWithTx(ctx context.Context, tx *sqlx.Tx, ids []int64) aerror.Error {
	delSql := `DELETE FROM product_image WHERE id IN (?)`

	delSql, args, err := sqlx.In(delSql, ids)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "del product_image failed")
	}

	return nil
}

func ProductImageUpdate(ctx context.Context, product *charon.ProductImage) aerror.Error {
	sqlStr := `UPDATE product_image
		SET product_id = :product_id,
		product_level = :product_level,
		uploader_type = :uploader_type,
		url = :url
		WHERE id = :id
`
	_, err := charon2.CharonDB.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func ProductImageUpdateWithTx(ctx context.Context, tx *sqlx.Tx, product *charon.ProductImage) aerror.Error {
	sqlStr := `UPDATE product_image
		SET product_id = :product_id,
		product_level = :product_level,
		uploader_type = :uploader_type,
		url = :url
		WHERE id = :id
`
	_, err := tx.NamedExecContext(ctx, sqlStr, product)
	if err != nil {
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "update failed")
	}

	return nil
}

func ProductImageSearch(ctx context.Context, query string, offset, limit int64) (int64, []*charon.ProductImage, aerror.Error) {
	formSql := `FROM product_image `
	whereSql := `WHERE uploader_type LIKE ? OR url LIKE ? `
	searchSelect := `SELECT id, product_id, product_level, uploader_type, url `
	countSelect := `SELECT count(*) `
	query = query + "%"

	countSelectSql := countSelect + formSql + whereSql
	var total int64
	err := charon2.CharonDB.QueryRowx(countSelectSql, query, query).Scan(&total)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select product_image failed")
	}

	if total == 0 {
		return 0, nil, nil
	}

	searchSelectSql := searchSelect + formSql + whereSql + "limit ?, ?"
	var images []*charon.ProductImage
	err = charon2.CharonDB.SelectContext(ctx, &images, searchSelectSql, query, query, offset, limit)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "select product_image failed")
	}

	return total, images, nil
}

func ProductImageSearchV2(ctx context.Context, in *proto.SubProductListRequest, filter *charonservice.SearchFilter) (int64, []*charon.ProductImage, aerror.Error) {
	wherePlaceholder := `{{whereCondition}}`
	sqlStr := fmt.Sprintf(`SELECT id, product_id, product_level, uploader_type, url FROM product_image %s limit ?, ?`, wherePlaceholder)
	countStr := fmt.Sprintf(`SELECT count(*) FROM product_image %s`, wherePlaceholder)

	var whereConditions []string
	var params []interface{}
	if len(filter.Id) > 0 {
		whereConditions = append(whereConditions, `id in (?)`)
		params = append(params, filter.Id)
	}

	if len(filter.Name) > 0 {
		whereConditions = append(whereConditions, "name uploader_type in (?)")
		params = append(params, filter.Name)
	}

	if len(filter.Desc) != 0 {
		whereConditions = append(whereConditions, "url in (?)")
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

	var images []*charon.ProductImage
	sqlStr, args, err = sqlx.In(sqlStr, params...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err = tx.SelectContext(ctx, &images, sqlStr, args...)
	if err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "select total category list failed")
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql commit error")
	}

	return count, images, nil
}
