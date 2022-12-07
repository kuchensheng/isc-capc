package category

import (
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

func (m *IscCapcCategory) GetCapcTableName() string {
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

//Create 新增分组信息
func (m *IscCapcCategory) Create(handler func() (bool, error)) (bool, error) {
	result := connetor.Db.Table(m.GetCapcTableName()).Create(m)
	if e := result.Error; e != nil {
		log.Warn().Msgf("分组信息新增失败%v", e)
		return false, common.CATEGORY_REGISTER_EXCEPTION.Exception(e)
	}
	return handler()
}

//Update 根据Id修改分组信息
func (m *IscCapcCategory) Update() (bool, error) {
	if m.ID == 0 {
		return false, common.CATEGORY_ID_ISNULL.Exception(nil)
	}
	//只更新非零字段
	result := connetor.Db.Table(m.GetCapcTableName()).Updates(m)
	if result.Error != nil {
		log.Warn().Msgf("分组信息更新失败%v", result.Error)
		return false, common.CATEGORY_REGISTER_EXCEPTION.Exception(result.Error)
	}
	return true, nil
}

//Delete 根据Id删除分组信息
func (m *IscCapcCategory) Delete() (bool, error) {
	if m.ID == 0 {
		return false, common.CATEGORY_ID_ISNULL.Exception(nil)
	}
	result := connetor.Db.Table(m.GetCapcTableName()).Delete(m)
	if e := result.Error; e != nil {
		log.Warn().Msgf("无法删除分组信息，%v", e)
		return false, common.CATEGORY_DELETE_EXCEPTION.Exception(e)
	}
	if result.RowsAffected < 1 {
		log.Warn().Msg("虽然没报错，但是没删除数据")
	}
	return true, nil
}
