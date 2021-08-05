package err

import (
	"github.com/bytehello/gcc-zero/common/api/protobuf"
	"github.com/bytehello/gcc-zero/common/errorx"
	"google.golang.org/grpc/status"
)

func ConvertErrorx(err error) error {
	s := status.Convert(err)
	d := s.Details()[0]
	switch info := d.(type) {
	case *protobuf.BizError:
		return errorx.NewCodeError(info.ErrCode, info.ErrMsg)
	default:
		return errorx.DefaultCodeError("未知错误")
	}
}
