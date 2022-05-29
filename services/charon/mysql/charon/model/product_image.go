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

func ProductImageInsert(ctx context.Context, image *charon.ProductImage) (*charon.ProductImage, aerror.Error) {
	logger.Infof("image: %s", logger.ToJson(image))

	insertSql := `INSERT INTO product_image (product_id, product_level, uploader_type, url, title) VALUES (:product_id, :product_level, :uploader_type, :url, :title)`
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

func ProductImageInsertWithTx(ctx context.Context, tx *sqlx.Tx, image *charon.ProductImage) (*charon.ProductImage, aerror.Error) {
	logger.Infof("image: %s", logger.ToJson(image))

	insertSql := `INSERT INTO product_image (product_id, product_level, uploader_type, url, title) VALUES (:product_id, :product_level, :uploader_type, :url, :title)`
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

func ProductImagesInsertWithTx(ctx context.Context, tx *sqlx.Tx, images []*charon.ProductImage) aerror.Error {
	valuesPlaceholder := `{{valuePlaceholder}}`
	valuePlaceholder := `(?, ?, ?, ?)`
	insertSql := fmt.Sprintf("INSERT INTO product_image (product_id, product_level, uploader_type, url) VALUES %s", valuesPlaceholder)

	var values []interface{}
	var valuePlaceholders []string
	for _, image := range images {
		values = append(values, image.ProductId, image.ProductLevel, image.UploaderType, image.Url)
		valuePlaceholders = append(valuePlaceholders, valuePlaceholder)
	}

	insertSql = strings.Replace(insertSql, valuesPlaceholder, strings.Join(valuePlaceholders, ","), -1)
	result, err := tx.ExecContext(ctx, insertSql, values...)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	_, err = result.RowsAffected()
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql rowsAffected failed")
	}

	return nil
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

func ProductImagesQueryByIds(ctx context.Context, id []int64) ([]*charon.ProductImage, aerror.Error) {
	querySql := `SELECT id, product_id, product_level, uploader_type, url FROM product_image WHERE id in (?)`
	var images []*charon.ProductImage

	querySql, arg, e := sqlx.In(querySql, id)
	if e != nil {
		return nil, aerror.NewErrorf(e, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	err := charon2.CharonDB.SelectContext(ctx, &images, querySql, arg...)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return images, nil
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

func ProductImageQueryByProductIdAndLevelWithTx(ctx context.Context, productId int64, productLevel int, tx ...*sqlx.Tx) ([]*charon.ProductImage, aerror.Error) {
	querySql := `SELECT id, product_id, product_level, uploader_type, url
				FROM product_image
				WHERE product_id = ? AND product_level = ?`
	var images []*charon.ProductImage
	var err error
	if tx == nil {
		err = charon2.CharonDB.SelectContext(ctx, &images, querySql, productId, productLevel)
	} else {
		err = tx[0].SelectContext(ctx, &images, querySql, productId, productLevel)
	}
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return images, nil
}

func ProductImageQueryByProductIdsAndLevelWithTx(ctx context.Context, productIds []int64, productLevel int, tx ...*sqlx.Tx) ([]*charon.ProductImage, aerror.Error) {
	querySql := `SELECT id, product_id, product_level, uploader_type, url
				FROM product_image
				WHERE product_level = ? AND product_id in (?)`
	var images []*charon.ProductImage
	var err error
	querySql, args, err := sqlx.In(querySql, productLevel, productIds)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in generate failed")
	}

	if tx == nil {
		err = charon2.CharonDB.SelectContext(ctx, &images, querySql, args...)
	} else {
		err = tx[0].SelectContext(ctx, &images, querySql, args...)
	}
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql execute error")
	}

	return images, nil
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

func ProductImageDeleteBySubProductIdAndLevelWithTx(ctx context.Context, subProductIds []int64, tx *sqlx.Tx) aerror.Error {
	if len(subProductIds) == 0 {
		return nil
	}
	delSql := `DELETE FROM product_image
			WHERE product_level = 2 AND product_id in (?)`

	delSql, args, err := sqlx.In(delSql, subProductIds)
	if err != nil {
		return aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "sql in failed")
	}

	_, err = tx.ExecContext(ctx, delSql, args...)
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
