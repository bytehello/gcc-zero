package logic

import (
	"context"
	"fmt"
	"github.com/bytehello/gcc-zero/internal/bizerror"

	"github.com/bytehello/gcc-zero/service/agent/cmd/rpc/agent"
	"github.com/bytehello/gcc-zero/service/agent/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReleaseKvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReleaseKvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseKvLogic {
	return &ReleaseKvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReleaseKvLogic) ReleaseKv(in *agent.ReleaseKvReq) (*agent.ReleaseKvReply, error) {
	err := l.svcCtx.ReleaseSvr.FilePutContents(in.FilePath, in.Value)
	if err != nil {
		return nil, bizerror.Customized(bizerror.ErrCodeReleaseKvFail, fmt.Sprintf("file put contents err: %s", err.Error()))
	}
	return &agent.ReleaseKvReply{}, nil
}
