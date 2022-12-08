package api

import (
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/rs/zerolog/log"
)

type ApiOperationRepository struct {
	Api        *IscCapcApiInfo
	Repository *apiParameterRepository
	Parameter  *IscCapcApiReqResp
}

//Create 创建API信息
func (op *ApiOperationRepository) Create() (bool, error) {
	tx := connetor.Db.Begin()
	defer func() {
		if x := recover(); x != nil {
			tx.Rollback()
		}
	}()
	log.Info().Msgf("创建API基本信息...")
	if ok, err := op.Api.Create(); !ok || err != nil {
		tx.Rollback()
		return ok, err
	}
	if op.Parameter != nil {
		log.Info().Msgf("保存或更新API参数信息")
		if ok, err := op.saveOrUpdateParameter(); !ok || err != nil {
			tx.Rollback()
			return ok, err
		}
	}
	tx.Commit()
	return true, nil
}

//Update 更新API信息
func (op *ApiOperationRepository) Update() (bool, error) {
	tx := connetor.Db
	defer func() {
		if x := recover(); x != nil {
			tx.Rollback()
		}
	}()
	log.Info().Msgf("更新API基本信息...")
	if ok, err := op.Api.Update(); !ok || err != nil {
		tx.Rollback()
		return ok, err
	}
	log.Info().Msgf("保存或更新API参数信息...")
	if op.Parameter != nil {
		if ok, err := op.saveOrUpdateParameter(); !ok || err != nil {
			tx.Rollback()
			return ok, err
		}
	}
	tx.Commit()
	return true, nil
}

//Delete 删除API信息
func (op *ApiOperationRepository) Delete() (bool, error) {
	tx := connetor.Db
	defer func() {
		if x := recover(); x != nil {
			tx.Rollback()
		}
	}()
	log.Info().Msgf("删除API信息...")
	if ok, err := op.Api.Delete(); !ok || err != nil {
		tx.Rollback()
		return ok, err
	}
	if op.Parameter != nil {
		log.Info().Msgf("删除API参数信息")
		if ok, err := op.Parameter.Delete(); !ok || err != nil {
			tx.Rollback()
			return ok, err
		}
	}

	tx.Commit()
	return true, nil
}

func (op *ApiOperationRepository) saveOrUpdateParameter() (bool, error) {
	if op.Parameter == nil {
		log.Info().Msg("创建api信息时，不含参数，无需做参数插入或更新处理")
		return true, nil
	}
	if one := op.Repository.GetOne(op.Parameter.ApiId, op.Parameter.Code); one != nil {
		//更新api参数信息
		one.Parameters = op.Parameter.Parameters
		one.Responses = op.Parameter.Responses
		one.Type = op.Parameter.Type
		if ok, err := one.Update(); !ok || err != nil {
			return ok, err
		} else {
			return true, nil
		}
	} else {
		return op.Parameter.Create()
	}
}
