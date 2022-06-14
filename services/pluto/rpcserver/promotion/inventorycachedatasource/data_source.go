package inventorycachedatasource

import (
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type dataSource string

const LocalVariableSource dataSource = "LocalVariableSource"
const MysqlDBSource dataSource = "MysqlDBSource"
const APISource dataSource = "APISource"

type DataSource interface {
	GetInventories(ctx context.Context) (map[int64]int64, aerror.Error)
	GetInventoriesByID(ctx context.Context, bookID int64) (map[int64]int64, aerror.Error)
}

func NewDataSource(sourceType dataSource) (DataSource, aerror.Error) {
	switch sourceType {
	case LocalVariableSource:
		return NewLocalVariablesSource(), nil
	case MysqlDBSource:
		return NewRedisDataSource(), nil
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "unknown sourceType: %s", sourceType)
	}
}
