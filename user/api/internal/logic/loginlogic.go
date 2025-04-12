package logic

import (
	"context"
	"img_analyse/user/commn/jwt"
	"img_analyse/user/rpc/types/user"

	"img_analyse/user/api/internal/svc"
	"img_analyse/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	rpcResponse, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &user.UserLoginRequest{
		Username:  req.UserName,
		Password1: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		Username: req.UserName,
		Password: req.Password,
		Role:     1,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		UserName: rpcResponse.Username,
		UserId:   int64(rpcResponse.UserId),
		Token:    token,
	}, nil
}
