package connetor

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type mysqlConnector struct {
}

func InitMysqlConnector() {
	connector := mysqlConnector{}
	e := connector.Open()
	println(e)
}
func (connector mysqlConnector) Open() error {
	db, err := gorm.Open(mysql.Open("isyscore:Isysc0re@tcp(10.30.30.95:23306)/isc_ecology_orchestration?parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	db = db.Debug()
	Db = db
	return nil
}
