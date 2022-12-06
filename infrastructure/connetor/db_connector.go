package connetor

import (
	"gorm.io/gorm"
)

var Db *gorm.DB

type Connector interface {
	//Open 打开DB连接，并保持长链
	Open() error
}
