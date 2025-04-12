package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"img_analyse/user/model"
	"img_analyse/user/rpc/internal/svc"
	"img_analyse/user/rpc/types/user"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	// 1. 参数校验
	if len(strings.TrimSpace(in.Username)) == 0 {
		return &user.UserLoginResponse{
			Msg: "用户名不能为空",
		}, errors.New("用户名不能为空")
	}

	if len(in.Password1) == 0 {
		return &user.UserLoginResponse{
			Msg: "密码不能为空",
		}, errors.New("密码不能为空")
	}

	// 2. 查询用户
	var u model.User
	err := l.svcCtx.DB.Where("username = ?", in.Username).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		l.Logger.Error("查询用户失败", logx.Field("err", err))
		return nil, errors.New("系统错误")
	}

	// 3. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password1))
	if err != nil {
		return &user.UserLoginResponse{
			Msg: "用户名或密码错误",
		}, errors.New("用户名或密码错误")
	}

	// 6. 返回响应
	return &user.UserLoginResponse{
		UserId:   uint32(u.ID),
		Username: u.Username,
	}, nil
}
