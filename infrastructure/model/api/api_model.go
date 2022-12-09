package api

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/infrastructure/vo/api"
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

func (model *IscCapcApiInfo) GetTableName() string {
	return tableName
}

func (model *IscCapcApiInfo) BeforeCreate(tx *gorm.DB) error {
	if err := model.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}

	model.Version = _VERSION
	return model.Valid()
}

func (model *IscCapcApiInfo) Create(ctx context.Context, db *gorm.DB) (bool, error) {
	if db == nil {
		db = connetor.GetDBWithTable(ctx, model.GetTableName())
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
func (model *IscCapcApiInfo) Update(ctx context.Context, db *gorm.DB) (bool, error) {
	if model.ID == 0 {
		return false, common.CATEGORY_ID_ISNULL.Exception(nil)
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

func (model *IscCapcApiInfo) Delete(ctx context.Context, db *gorm.DB) (bool, error) {
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
	if result.RowsAffected < 1 {
		log.Warn().Msgf("api信息不存在")
		return false, common.API_NOT_EXISTS.Exception(nil)
	}
	log.Info().Msgf("信息删除成功,ID=%d", model.ID)
	return true, nil
}

func (model *IscCapcApiInfo) DeleteBatch(ctx context.Context, db *gorm.DB, search api.SearchVO) (bool, error) {
	if db == nil {
		db = connetor.GetDBWithTable(ctx, model.GetTableName())
	}
	condition := struct {
		Summary string
		Path    string
		Type    []int
		Id      []int
		Code    []string
		Method  []string
	}{
		search.Name, search.Path, search.Types, search.Ids, search.Codes, search.Methods,
	}
	result := db.Delete(condition)
	return result.Error == nil, result.Error
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
