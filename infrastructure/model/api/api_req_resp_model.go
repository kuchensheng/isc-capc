package api

import (
	"fmt"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/model"
	"gorm.io/gorm"
	"strings"
)

const _tableName = "isc_capc_api_req_resp"

type IscCapcApiReqResp struct {
	model.BaseModel

	ApiId      int    `json:"api_id"`
	Parameters string `json:"parameters"`
	Responses  string `json:"responses"`
	Type       int    `json:"type"`
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
