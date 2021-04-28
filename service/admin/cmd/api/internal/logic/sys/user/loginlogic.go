package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/service/admin/model"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
	adminUser, err := l.svcCtx.AdminUser.FindOneByUsername(req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.DefaultCodeError("用户名不存在")
	default:
		return nil, errorx.DefaultCodeError(err.Error())
	}
	if adminUser.Password != req.Password {
		return nil, errorx.DefaultCodeError("密码错误")
	}
	// jwt
	var token string
	accessExpire := l.svcCtx.Config.Jwt.AccessExpire
	token, err = l.getJwtToken(adminUser.Id, l.svcCtx.Config.Jwt.AccessExpire, l.svcCtx.Config.Jwt.AccessSecret, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	return &types.LoginReply{
		Id:           adminUser.Id,
		Username:     adminUser.UserName,
		AccessToken:  token,
		AccessExpire: accessExpire,
	}, nil
}

func (l *LoginLogic) getJwtToken(uid int64, seconds int64, secret string, iat int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["uid"] = uid
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
