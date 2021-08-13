package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/internal/bizerror"
	"os"

	"github.com/bytehello/gcc-zero/service/agent/cmd/rpc/agent"
	"github.com/bytehello/gcc-zero/service/agent/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReleaseKvStatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReleaseKvStatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseKvStatLogic {
	return &ReleaseKvStatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReleaseKvStatLogic) ReleaseKvStat(in *agent.ReleaseKvStatReq) (*agent.ReleaseKvStatReply, error) {
	var (
		stat os.FileInfo
		err  error
	)
	if stat, err = l.svcCtx.ReleaseSvr.ConfigFileStat(in.FilePath); err != nil {
		return nil, bizerror.New(bizerror.ErrCodeReleaseKvStat)
	}
	return &agent.ReleaseKvStatReply{
		Name:    stat.Name(),
		Size:    stat.Size(),
		ModTime: stat.ModTime().Format("2006-01-02 15:04:05"),
	}, nil
}
