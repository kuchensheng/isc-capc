package connetor

import (
	"fmt"
	"github.com/isyscore/isc-gobase/config"
	"github.com/isyscore/isc-gobase/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type mysqlConnector struct {
}

func init() {
	connector := mysqlConnector{}
	if e := connector.Open(); e != nil {
		fmt.Println(e)
	}
}

const _DB = "%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FShanghai"

func (connector mysqlConnector) Open() error {
	userName := config.GetValueStringDefault("base.datasource.userName", "isyscore")
	password := config.GetValueStringDefault("base.datasource.password", "Isysc0re")
	host := config.GetValueString("base.datasource.host")
	port := config.GetValueIntDefault("base.datasource.port", 3306)
	database := config.GetValueString("base.datasource.database")
	url := fmt.Sprintf(_DB, userName, password, host, port, database)
	logger.Info("连接数据库:%s", url)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		logger.Fatal("无法打开连接：%v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	db = db.Debug()
	Db = db
	return nil
}
