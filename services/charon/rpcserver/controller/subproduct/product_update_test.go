package subproduct

import (
	"amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestHandlerUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ....")
	}

	charon.Mock()

	in := &proto.SubProductUpdateRequest{
		Id:          9,
		Name:        "new name",
		Desc:        "new desc",
		ProductId:   102,
		Currency:    "CNY",
		Price:       10001,
		Stock:       10002,
		AttributeId: []int64{4, 5, 6},
	}

	resp, err := HandlerUpdate(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("resp: %s", logger.ToJson(resp))
}
