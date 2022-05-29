package subproduct

import (
	"amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon/model"
	"amusingx.fit/amusingx/services/charon/uploader"
	"context"
)

func ProductPictures(ctx context.Context, req *proto.SubProductPicturesRequest) (*proto.SubProductPicturesResponse, error) {
	pictures, err := model.ProductImageQueryByProductIdsAndLevelWithTx(ctx, req.SubProductIds, charon.ProductLevelSubProduct)
	if err != nil {
		return nil, err
	}

	subProductIDPictureMap := make(map[int64][]*proto.Picture)
	for _, p := range pictures {
		uploader, err := uploader.NewUploader(p.UploaderType)
		if err != nil {
			return nil, err
		}

		image, err := uploader.GetPictureInfo(ctx, p.Id)
		if err != nil {
			return nil, err
		}

		picture := &proto.Picture{
			RawFile:    nil,
			Src:        image.Url,
			Title:      p.Title,
			Id:         p.Id,
			Url:        image.Url,
			UploadType: int64(p.UploaderType),
		}

		if _, ok := subProductIDPictureMap[p.ProductId]; !ok {
			subProductIDPictureMap[p.ProductId] = []*proto.Picture{picture}
		} else {
			subProductIDPictureMap[p.ProductId] = append(subProductIDPictureMap[p.ProductId], picture)
		}
	}

	var subProductPictures []*proto.SubProductPicture
	for subProductId, pictures := range subProductIDPictureMap {
		subProductPictures = append(subProductPictures, &proto.SubProductPicture{
			SubProductId: subProductId,
			Pictures:     pictures,
		})
	}

	return &proto.SubProductPicturesResponse{SubProductPictures: subProductPictures}, nil
}
