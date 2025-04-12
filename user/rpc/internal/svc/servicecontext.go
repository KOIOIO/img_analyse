package svc

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"img_analyse/user/commn/Logger"
	"img_analyse/user/commn/init_gorm"
	"img_analyse/user/model"
	"img_analyse/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		Logger.Log.Error("数据库迁移失败", zap.Error(err))
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
