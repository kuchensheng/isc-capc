package api

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type apiParameterRepository struct {
	model.BaseRepository
}

var ApiParameterRepository = &apiParameterRepository{model.BaseRepository{DB: connetor.Db.Table(_tableName)}}

func (repository *apiParameterRepository) GetDB(context context.Context) *gorm.DB {
	return repository.BaseRepository.GetDB(context)
}

func (repository *apiParameterRepository) GetOne(apiId int, code string, context context.Context) *IscCapcApiReqResp {
	if apiId == 0 && code == "" {
		return nil
	}
	db := repository.GetDB(context)
	search := &struct {
		ApiId int
		Code  string
	}{apiId, code}
	parameter := NewIscCapcApiReqResp()
	result := db.Where(search).Take(parameter)
	if e := result.Error; e != nil {
		//查询异常
		log.Warn().Msgf("无法获取apiReqResp信息,%v", e)
		return nil
	}
	return parameter
}
