package model

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"gorm.io/gorm"
)

//Repository 基础仓库
type Repository interface {
	GetDB() *gorm.DB
}

type BaseRepository struct {
}

func (repository *BaseRepository) GetDB(context context.Context) *gorm.DB {
	return connetor.GetDB(context)
}
