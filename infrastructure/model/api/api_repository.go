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

func (repository *apiRepository) GetOne(vo api.DetailVO) (IscCapcApiInfo, bool) {
	db := repository.buildDetailDB(vo)
	result := IscCapcApiInfo{}
	db = db.Take(&result)
	if db.Error == nil && db.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (repository *apiRepository) buildDetailDB(vo api.DetailVO) *gorm.DB {
	db := repository.GetDB()
	db.Where(&vo)
	return db
}

func (repository *apiRepository) GetBaseApiList() {

}
