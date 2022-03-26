package comm

import (
	"amusingx.fit/amusingx/protos/comm/response"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Response(data proto.Message, e aerror.Error) *response.CommResponse {
	resp := &response.CommResponse{}
	if e != nil {
		resp.Succeed = false
		resp.Error = &response.Error{
			Code:    e.Code(),
			Message: e.Message(),
		}

		return resp
	}

	any, err := anypb.New(data)
	if err != nil {
		resp.Succeed = false
		resp.Error = &response.Error{
			Code:    aerror.Code.BObjectToAnyFailed,
			Message: "unexpected data",
		}
	}

	return &response.CommResponse{Data: any, Succeed: true}
}
