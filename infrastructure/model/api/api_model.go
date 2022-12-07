package api

import (
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"math/rand"
	"strings"
)

const tableName = "isc_capc_api_info"

type IscCapcApiInfo struct {
	model.BaseModel

	//Summary api名称
	Summary string `json:"summary"`
	//Path api路径
	Path string `json:"path"`
	//Method 请求方法,GET、POST、DELETE、PUT、PATCH等
	Method string `json:"method"`
	//Code api唯一标识
	Code string `json:"code"`
	//Type api类型，OS、NATIVE、POLYMERIC、LIGHT、UDMP、TDDM等
	Type int `json:"type"`
	//CategoryId 分组ID
	CategoryId int `json:"group_id" gorm:"column:group_id"`
	//CategoryPath 分组全路径ID
	CategoryPath string `json:"group_path" gorm:"column:group_path"`
	//CategoryFullName 分组全路径名称
	CategoryFullName string `json:"group_full_name" gorm:"column:group_full_name"`
	//Status 生命周期状态
	Status int `json:"status"`
	//Tags 标签信息
	Tags string `json:"tags"`
	//Version 版本信息，默认1.0.0
	Version string `json:"version"`
	//Description 描述信息
	Description string `json:"description"`
	//Protocol 协议信息，HTTP、HTTPS、WS、WSS、WQ等，默认HTTP
	Protocol string `json:"protocol"`
	//Principal 责任人信息
	Principal string `json:"principal"`
	//AuthType 鉴权方式
	AuthType int `json:"auth_type"`
	//AuthConfig 鉴权配置
	AuthConfig string `json:"auth_config"`
	//Cosumes 解析方式，是个数组
	Consumes string `json:"consumes"`
}

func NewIscCapcApiInfo() *IscCapcApiInfo {
	return &IscCapcApiInfo{
		Method:   "GET",
		Code:     randString(16),
		Type:     1,
		Status:   0,
		Version:  _VERSION,
		Protocol: "HTTP",
		Consumes: `["application/json"]`,
	}
}

func (model *IscCapcApiInfo) Valid() error {
	if len(model.Code) == 0 {
		return common.API_CODE_IS_NULL.Exception(nil)
	}
	if len(model.Path) == 0 {
		return common.API_PATH_IS_NULL.Exception(nil)
	}
	if len(model.Summary) == 0 {
		return common.API_NAME_IS_NULL.Exception(nil)
	}
	if model.CategoryId == 0 {
		return common.API_CATEGORY_IS_NULL.Exception(nil)
	}
	return nil
}

func (model *IscCapcApiInfo) GetCapcTableName() string {
	return tableName
}

func (model *IscCapcApiInfo) BeforeCreate(tx *gorm.DB) error {
	if err := model.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}

	return model.Valid()
}

var _VERSION = "1.0.0"
var _CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func randString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(_CHARS[rand.Intn(length)])
	}
	return str.String()
}

func (model *IscCapcApiInfo) Create() error {
	result := connetor.Db.Table(model.GetCapcTableName()).Create(model)
	if err := result.Error; err != nil {
		log.Warn().Msgf("api信息新增失败,%v", err)
		return common.API_REGISTER_EXCEPTION.Exception(err.Error())
	} else if result.RowsAffected < 1 {
		log.Warn().Msgf("未知原因导致没有新增成功")
		return common.UNKNOWN_EXCEPTION.Exception(nil)
	}
	return nil
}

func (model *IscCapcApiInfo) Update() error {
	result := connetor.Db.Table(model.GetCapcTableName()).Updates(model)
	if err := result.Error; err != nil {
		log.Warn().Msgf("api信息更新失败,%v", err)
		return common.API_REGISTER_EXCEPTION.Exception(err.Error())
	} else if result.RowsAffected < 1 {
		log.Warn().Msgf("未知原因导致没有更新成功")
		return common.UNKNOWN_EXCEPTION.Exception(nil)
	}
	return nil
}

func (model *IscCapcApiInfo) Delete() error {
	result := connetor.Db.Table(model.GetCapcTableName()).Delete(model)
	if err := result.Error; err != nil {
		log.Warn().Msgf("api信息删除失败,%v", err)
		return common.API_REGISTER_EXCEPTION.Exception(err.Error())
	} else if result.RowsAffected < 1 {
		log.Warn().Msgf("未知原因导致没有删除成功")
		return common.UNKNOWN_EXCEPTION.Exception(nil)
	}
	return nil
}
