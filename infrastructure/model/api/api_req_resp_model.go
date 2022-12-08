package api

import (
	"fmt"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"gorm.io/gorm"
	"strings"
)

const _tableName = "isc_capc_api_req_resp"

type IscCapcApiReqResp struct {
	model.BaseModel

	//关联的ApiId
	//fixme 未来将被废弃
	ApiId int `json:"api_id"`
	//关联的ApiCode
	Code string `json:"code"`
	//Parameters 入参信息模型
	Parameters string `json:"parameters"`
	//Resposes 出参信息模型
	Responses string `json:"responses"`
	//Type 出参类型，JSON/XML
	Type int `json:"type"`
}

func NewIscCapcApiReqResp() *IscCapcApiReqResp {
	return &IscCapcApiReqResp{
		Type: 0,
	}
}

func (model *IscCapcApiReqResp) Valid() error {
	if model.ApiId == 0 {
		return common.API_P_APIID_IS_NULL.Exception(nil)
	}
	if model.Parameters == "" {
		model.Parameters = "[]"
	} else if !strings.HasPrefix(model.Parameters, "[") && !strings.HasSuffix(model.Parameters, "]") {
		model.Parameters = fmt.Sprintf("[%s]", model.Parameters)
	}
	if model.Responses == "" {
		model.Responses = "{}"
	}
	return nil
}

func (model *IscCapcApiReqResp) GetCapcTableName() string {
	return _tableName
}

func (model *IscCapcApiReqResp) BeforeCreate(tx *gorm.DB) error {
	if err := model.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}
	return model.Valid()
}

func (model *IscCapcApiReqResp) BeforeUpdate(tx *gorm.DB) error {
	return model.BaseModel.BeforeUpdate(tx)
}

func (model *IscCapcApiReqResp) Delete() (bool, error) {
	if model.ApiId == 0 && model.Code == "" {
		return false, common.BAD_REQUEST.Exception("apiId或code不能同时为空")
	}
	db := connetor.Db.Table(model.GetCapcTableName())
	deleteParam := &struct {
		ApiId int
		Code  string
	}{
		model.ApiId,
		model.Code,
	}
	result := db.Delete(deleteParam)
	if err := result.Error; err != nil {
		return false, common.DELETE_EXCEPTION.Exception(err.Error())
	}
	return true, nil
}
