package api

import (
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/infrastructure/vo/api"
	"gorm.io/gorm"
)

type apiRepository struct {
	model.BaseRepository
}

var ApiRepository = &apiRepository{model.BaseRepository{DB: connetor.Db.Table(tableName)}}

func (repository *apiRepository) GetDB() *gorm.DB {
	return repository.BaseRepository.GetDB()
}

func (repository *apiRepository) GetOne(vo api.SearchVO) (IscCapcApiInfo, bool) {
	db := repository.GetDB()
	if vo.Code != "" {
		db = db.Where("code = ?", vo.Code)
	}
	result := IscCapcApiInfo{}
	db = db.Take(&result)
	if db.Error == nil && db.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (repository *apiRepository) GetBaseApiList() {

}
