package category

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/util"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

const tableName = "isc_capc_category"

type IscCapcCategory struct {
	model.BaseModel `gorm:"embedded"`

	//Name 分组名称
	Name string `json:"name"`
	//Type 分组类型
	Type int `json:"type"`
	//ParentId上级ID，为空，表示首层
	ParentId int `json:"parent_id"`
	//ParentCode 上级唯一标识，
	ParentCode string `json:"parent_code"`
	//Introduce 应用或分组描述
	Introduce string `json:"introduce"`
	//Code 当前分组Code
	Code string `json:"code"`
	//Tags 名称路径，将被用于全路径识别
	Tags string `json:"tags"`
	//CodePath code路径，将被用于全路径识别
	CodePath string `json:"code_path"`
	//Icon 应用或分组图标
	Icon string `json:"icon"`
	//Config 公共配置信息
	Config string `json:"config"`
	//Version 版本号，默认1.0
	Version string `json:"version"`
	//Remark 备注
	Remark string `json:"remark"`
}

func NewIscCapcCategory() *IscCapcCategory {
	return &IscCapcCategory{}
}

func (m *IscCapcCategory) GetTableName() string {
	return tableName
}

func (m *IscCapcCategory) BeforeCreate(tx *gorm.DB) error {
	if e := m.BaseModel.BeforeCreate(tx); e != nil {
		return e
	} else {
		if m.Code == "" {
			m.Code = util.RandString(16)
		}
		if m.Type == 0 {
			m.Type = 1
		}
	}
	return nil
}

func (model *IscCapcCategory) Create(ctx context.Context, db *gorm.DB) (bool, error) {
	if db == nil {
		db = connetor.GetDBWithTable(ctx, model.GetTableName())
	}
	tenantId := ctx.Value(common.TENANTID)
	if tenantId != nil {
		model.TenantId = tenantId.(string)
	}
	result := db.Create(model)
	if e := result.Error; e != nil {
		log.Warn().Msgf("信息注册异常,%v", e)
		return false, common.REGISTER_EXCEPTION.Exception(e.Error())
	}
	log.Info().Msgf("信息注册成功,ID=%d", model.ID)
	return true, nil
}

//Update 根据Id修改分组信息
func (model *IscCapcCategory) Update(ctx context.Context, db *gorm.DB) (bool, error) {
	if model.ID == 0 {
		return false, common.ID_IS_NULL.Exception(nil)
	}
	if db == nil {
		db = connetor.GetDBWithTable(ctx, model.GetTableName())
	}
	//只更新非零字段
	result := db.Updates(model)
	if result.Error != nil {
		log.Warn().Msgf("信息更新异常,%v", result.Error)
		return false, common.UPDATE_EXCEPTION.Exception(result.Error.Error())
	}
	log.Info().Msgf("信息更新成功,ID=%d", model.ID)
	return true, nil
}

func (model *IscCapcCategory) Delete(ctx context.Context, db *gorm.DB) (bool, error) {
	if model.ID == 0 {
		return false, common.ID_IS_NULL.Exception(nil)
	}
	if db == nil {
		db = connetor.GetDBWithTable(ctx, model.GetTableName())
	}
	db.Where("id = ?", model.ID)
	result := db.Delete(model)
	if e := result.Error; e != nil {
		log.Warn().Msgf("信息删除异常,%v", e)
		return false, common.DELETE_EXCEPTION.Exception(e.Error())
	}
	log.Info().Msgf("信息删除成功,ID=%d", model.ID)
	return true, nil
}
