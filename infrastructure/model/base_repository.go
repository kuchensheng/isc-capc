package model

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"gorm.io/gorm"
)

//Repository 基础仓库
type Repository interface {
	GetDB() *gorm.DB
}

type BaseRepository struct {
	DB *gorm.DB
}

func (repository *BaseRepository) GetDB(context context.Context) *gorm.DB {
	db := repository.DB
	if db == nil {
		db = connetor.Db
		repository.DB = db
	}
	db = db.WithContext(context).Where("tenant_id = ?", context.Value(common.TENANTID))
	return db
}
