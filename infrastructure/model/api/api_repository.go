package api

import (
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"gorm.io/gorm"
)

type apiRepository struct {
	model.BaseRepository
}

var ApiRepository *apiRepository = &apiRepository{model.BaseRepository{DB: connetor.Db, TableName: tableName}}

func (repository *apiRepository) GetDB() *gorm.DB {
	return repository.BaseRepository.GetDB()
}

func (respository *apiRepository) GetBaseApiList() {

}
