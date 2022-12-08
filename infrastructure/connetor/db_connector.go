package connetor

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/common"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

type Connector interface {
	//Open 打开DB连接，并保持长链
	Open() error
}

func GetDB(ctx context.Context) *gorm.DB {
	return gormDB.WithContext(ctx).Where("tenant_id = ?", ctx.Value(common.TENANTID))
}
