package api

import (
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"gorm.io/gorm"
)

type apiParameterRepository struct {
	model.BaseRepository
}

var ApiParameterRepository = &apiParameterRepository{model.BaseRepository{DB: connetor.Db.Table(_tableName)}}

func (repository *apiParameterRepository) GetDB() *gorm.DB {
	return repository.BaseRepository.GetDB()
}
