package inventorycachedatasource

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

const LocalVariableSource = "LocalVariableSource"
const MysqlDBSource = "MysqlDBSource"

type DataSource interface {
	GetInventories(ctx context.Context) (map[int64]int64, *xerror.Error)
	GetInventoriesByID(ctx context.Context, bookID int64) (map[int64]int64, *xerror.Error)
}

func NewDataSource(sourceType string) (DataSource, *xerror.Error) {
	switch sourceType {
	case LocalVariableSource:
		return NewLocalVariablesSource(), nil
	case MysqlDBSource:
		return NewRedisDataSource(), nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "unknown sourceType: %s", sourceType)
	}
}
