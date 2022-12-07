package category

import (
	"encoding/json"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/model/category"
	category2 "github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/rs/zerolog/log"
)

type CategoryDTO struct {
	ID int `json:"id"`
	//Name 分组名称
	Name string `json:"name"`
	//Type 分组类型
	Type category2.CategoryType `json:"type,omitempty"`
	//ParentId上级ID，为空，表示首层
	ParentId int `json:"parent_id"`
	//ParentCode 上级唯一标识，
	ParentCode string `json:"parent_code"`
	//Introduce 应用或分组描述
	Introduce string `json:"introduce"`
	//Code 当前分组Code,唯一标识
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

func (dto *CategoryDTO) Check() error {
	if len(dto.Name) == 0 {
		return common.CATEGORY_NAME_ISNULL.Exception(nil)
	}
	//if dto.Type == 0 {
	//	dto.Type = category.NATIVE
	//}
	//if dto.Version == "" {
	//	dto.Version = "1.0.0"
	//}
	return nil
}

func (dto CategoryDTO) Dto2Do() *category.IscCapcCategory {
	data, _ := json.Marshal(dto)
	model := &category.IscCapcCategory{}
	err := json.Unmarshal(data, model)
	if err != nil {
		log.Warn().Msgf("分组信息转换失败,%v", err)
		return nil
	}
	return model
}

func Do2DTO(model category.IscCapcCategory) *CategoryDTO {
	data, _ := json.Marshal(model)
	dto := &CategoryDTO{}
	err := json.Unmarshal(data, dto)
	if err != nil {
		log.Warn().Msgf("分组信息转换失败,%v", err)
		return nil
	}
	return dto
}
