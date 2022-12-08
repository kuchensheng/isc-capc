package api

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/infrastructure/vo/api"
	"gorm.io/gorm"
)

type apiRepository struct {
	model.BaseRepository
}

var ApiRepository = &apiRepository{model.BaseRepository{DB: connetor.Db.Table(tableName)}}

func (repository *apiRepository) GetDB(context context.Context) *gorm.DB {
	return repository.BaseRepository.GetDB(context)
}

func (repository *apiRepository) GetOne(vo api.DetailVO, context context.Context) (IscCapcApiInfo, bool) {
	db := repository.buildDetailDB(vo, context)
	result := IscCapcApiInfo{}
	db = db.Take(&result)
	if db.Error == nil && db.RowsAffected > 0 {
		return result, true
	}
	return result, false
}

func (repository *apiRepository) buildDetailDB(vo api.DetailVO, context context.Context) *gorm.DB {
	db := repository.GetDB(context)
	db.Where(&vo)
	return db
}

func (repository *apiRepository) GetBaseApiList(context context.Context) {

}
