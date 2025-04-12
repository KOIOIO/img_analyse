package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"img_analyse/user/model"
	"img_analyse/user/rpc/internal/svc"
	"img_analyse/user/rpc/types/user"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	// todo: add your logic here and delete this line

	// 1. 参数校验
	if len(strings.TrimSpace(in.Username)) < 4 {
		return &user.UserRegisterResponse{
			Msg: "用户名长度不能小于4位",
		}, errors.New("用户名长度不能小于4位")
	}

	if len(in.Password1) < 6 {
		return &user.UserRegisterResponse{
			Msg: "密码长度不能小于6位",
		}, errors.New("密码长度不能小于6位")
	}

	if in.Password1 != in.Password2 {
		return &user.UserRegisterResponse{
			Msg: "两次输入的密码不一致",
		}, errors.New("两次输入的密码不一致")
	}

	// 2. 检查用户名是否已存在
	var count int64
	err := l.svcCtx.DB.Model(&model.User{}).Where("username = ?", in.Username).Count(&count).Error
	if err != nil {
		l.Logger.Error("查询用户失败", logx.Field("err", err))
		return &user.UserRegisterResponse{
			Msg: "系统错误",
		}, errors.New("系统错误")
	}

	if count > 0 {
		return &user.UserRegisterResponse{
			Msg: "用户名已存在",
		}, errors.New("用户名已存在")
	}

	// 3. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password1), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Error("密码加密失败", logx.Field("err", err))
		return &user.UserRegisterResponse{
			Msg: "密码加密失败",
		}, errors.New("系统错误")
	}

	// 4. 创建用户
	newUser := model.User{
		Username: in.Username,
		Password: string(hashedPassword),
	}

	err = l.svcCtx.DB.Create(&newUser).Error
	if err != nil {
		l.Logger.Error("创建用户失败", logx.Field("err", err))
		return nil, errors.New("系统错误")
	}

	// 5. 返回响应
	return &user.UserRegisterResponse{
		UserId:   uint32(newUser.ID),
		Username: newUser.Username,
		Msg:      "注册成功",
	}, nil
}
