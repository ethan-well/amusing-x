package categoryapp

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func CreateCategory(ctx context.Context, category *charonservice.CategoryCreateRequest) (*charon.Category, aerror.Error) {
	c, err := model.QueryCategoryByName(ctx, category.Name)
	if err != nil {
		return nil, err
	}

	if c != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "category '%s' is already existed", category.Name)
	}

	c, err = model.InsetCategory(ctx, &charon.Category{
		Name: category.Name,
		Desc: category.Desc,
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}
