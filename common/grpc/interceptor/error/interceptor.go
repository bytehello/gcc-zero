package error

import (
	"context"
	"github.com/bytehello/gcc-zero/common/api/protobuf"
	"github.com/bytehello/gkit/bizerror"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (retResp interface{}, retErr error) {
	retResp, retErr = handler(ctx, req)
	if retErr == nil {
		return
	}
	if err, ok := retErr.(*bizerror.BizError); ok {
		logx.WithContext(ctx).Errorf("Biz Error: method: %s request: `%s` err: `%s`",
			info.FullMethod, req, err)
		s, _ := status.New(codes.FailedPrecondition, err.Msg).WithDetails(&protobuf.BizError{
			ErrCode: int64(err.Code),
			ErrMsg:  err.Msg,
		})
		retErr = s.Err()
		return
	}
	logx.WithContext(ctx).Errorf("Unknown Error: method: %s request: `%s` err: `%s`",
		info.FullMethod, req, retErr)
	retErr = status.Error(codes.Unknown, "Unknown Error")
	return retResp, retErr
}
