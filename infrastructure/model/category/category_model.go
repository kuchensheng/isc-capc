package category

import (
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/util"
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
