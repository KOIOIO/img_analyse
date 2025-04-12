package logic

import (
	"context"
	"errors"
	"img_analyse/user/rpc/types/user"

	"img_analyse/user/api/internal/svc"
	"img_analyse/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &user.UserRegisterRequest{
		Username:  req.UserName,
		Password1: req.Password1,
		Password2: req.Password2,
	})
	if err != nil {
		l.Logger.Error("注册失败", logx.Field("err", err))
		return nil, errors.New("注册失败")
	}
	return &types.RegisterResponse{
		UserId:   int64(response.UserId),
		UserName: response.Username,
	}, nil

}
