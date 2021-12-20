package categoryapp

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charons/service"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func CreateCategory(ctx context.Context, category *charonservice.CategoryCreateRequest) (*charon.Category, *xerror.Error) {
	c, err := model.InsetCategory(ctx, &charon.Category{
		Name: category.Name,
		Desc: category.Desc,
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}
