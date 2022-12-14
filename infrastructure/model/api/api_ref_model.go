package api

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/infrastructure/vo/api"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ApiOperationRepository struct {
	model.BaseRepository
	Api        *IscCapcApiInfo
	Repository *apiParameterRepository
	Parameter  *IscCapcApiReqResp
	SearchVo   api.SearchVO
}

func (repository *ApiOperationRepository) GetDB(context context.Context) *gorm.DB {
	return repository.BaseRepository.GetDB(context)
}

//Create 创建API信息
func (op *ApiOperationRepository) Create(ctx context.Context) (bool, error) {
	tx := op.GetDB(ctx).Begin()
	defer func() {
		if x := recover(); x != nil {
			log.Error().Msgf("无法创建API信息，%v", x)
			tx.Rollback()
		}
	}()
	log.Info().Msgf("创建API基本信息...")
	if ok, err := op.Api.Create(ctx, tx); !ok || err != nil {
		tx.Rollback()
		return ok, err
	}
	op.Parameter.ApiId = op.Api.ID
	op.Parameter.Code = op.Api.Code
	if op.Parameter != nil {
		log.Info().Msgf("保存或更新API参数信息")
		if ok, err := op.saveOrUpdateParameter(ctx, nil); !ok || err != nil {
			tx.Rollback()
			return ok, err
		}
	}
	tx.Commit()
	return true, nil
}

//Update 更新API信息
func (op *ApiOperationRepository) Update(ctx context.Context) (bool, error) {
	tx := op.GetDB(ctx).Begin()
	defer func() {
		if x := recover(); x != nil {
			tx.Rollback()
		}
	}()
	log.Info().Msgf("更新API基本信息...")
	if ok, err := op.Api.Update(ctx, tx); !ok || err != nil {
		tx.Rollback()
		return ok, err
	}
	log.Info().Msgf("保存或更新API参数信息...")
	op.Parameter.ApiId = op.Api.ID
	op.Parameter.Code = op.Api.Code
	if op.Parameter != nil {
		if ok, err := op.saveOrUpdateParameter(ctx, tx); !ok || err != nil {
			tx.Rollback()
			return ok, err
		}
	}
	tx.Commit()
	return true, nil
}

//Delete 删除API信息
func (op *ApiOperationRepository) Delete(ctx context.Context) (bool, error) {
	tx := op.GetDB(ctx).Begin()
	defer func() {
		if x := recover(); x != nil {
			tx.Rollback()
		}
	}()
	log.Info().Msgf("删除API信息...")
	if ok, err := op.Api.Delete(ctx, tx); !ok || err != nil {
		tx.Rollback()
		return ok, err
	}
	op.Parameter.ApiId = op.Api.ID
	op.Parameter.Code = op.Api.Code
	if op.Parameter != nil {
		log.Info().Msgf("删除API参数信息")
		if ok, err := op.Parameter.Delete(ctx, tx); !ok || err != nil {
			tx.Rollback()
			return ok, err
		}
	}

	tx.Commit()
	return true, nil
}

//DeleteBatch 批量删除API信息
func (op *ApiOperationRepository) DeleteBatch(ctx context.Context) (bool, error) {
	tx := op.GetDB(ctx).Begin()
	defer func() {
		if x := recover(); x != nil {
			tx.Rollback()
		}
	}()
	log.Info().Msgf("删除API信息...")
	//todo 还未完成
	if op.SearchVo.Ids != nil {
		return op.Api.DeleteBatch(ctx, tx, op.SearchVo)
	}
	return true, nil
}

func (op *ApiOperationRepository) saveOrUpdateParameter(ctx context.Context, tx *gorm.DB) (bool, error) {
	if op.Parameter == nil {
		log.Info().Msg("创建api信息时，不含参数，无需做参数插入或更新处理")
		return true, nil
	}
	if one := op.Repository.GetOne(op.Parameter.ApiId, op.Parameter.Code, ctx); one != nil {
		//更新api参数信息
		one.Parameters = op.Parameter.Parameters
		one.Responses = op.Parameter.Responses
		one.Type = op.Parameter.Type
		if ok, err := one.Update(ctx, tx); !ok || err != nil {
			return ok, err
		} else {
			return true, nil
		}
	} else {
		return op.Parameter.Create(ctx, tx)
	}
}
