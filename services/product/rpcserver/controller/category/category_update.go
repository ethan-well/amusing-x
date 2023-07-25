package category

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/services/product/mysql/charon/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerUpdate(ctx context.Context, req *charonservice.CategoryUpdateRequest) (*charonservice.CategoryUpdateResponse, aerror.Error) {
	err := model.UpdateCategory(ctx, &charon.Category{
		ID:   req.Category.Id,
		Name: req.Category.Name,
		Desc: req.Category.Desc,
	})

	if err != nil {
		return nil, err
	}

	return &charonservice.CategoryUpdateResponse{Result: true}, nil
}
