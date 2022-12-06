package model

import (
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/util"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type BaseModel struct {
	//ID 主键
	ID int `json:"id"`
	//CreateTime 创建时间
	CreateTime string `json:"create_time"`
	//UpdateTime 最近更新时间
	UpdateTime string `json:"update_time"`
	//TenantId 租户Id信息
	TenantId string `json:"tenant_id"`
}

type GetTableNameInterface interface {
	GetCapcTableName() string
}

func (model *BaseModel) GetCapcTableName() string {
	return ""
}

func (model *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if model.ID != 0 {
		return common.NOT_ALLOW.Exception("创建时,id必须为空")
	}
	if model.TenantId == "" {
		return common.NOT_ALLOW.Exception("租户Id不能为空")
	}
	return nil
}

func (model *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	if model.TenantId == "" {
		return common.NOT_ALLOW.Exception("租户Id不能为空")
	}
	//检查信息是否存在
	result := tx.Table(model.GetCapcTableName()).First(model)
	if result.RowsAffected < 0 {
		return common.CATEGORY_NOT_EXISTS.Exception(nil)
	}
	return nil
}

//GetAllFieldAndValues 获取当前的对象所有的字段名和字段值
func (model *BaseModel) GetAllFieldAndValues() []common.Pair[string, any] {
	of := reflect.ValueOf(model)
	typeOf := of.Type()
	var result []common.Pair[string, any]
	for i := 0; i < of.NumField(); i++ {
		f := of.Field(i)
		result = append(result, common.CreatePair(typeOf.Field(i).Name, f.Interface()))
	}
	return result
}

//GetNotNullFieldAndValues 获取当前的对象所有的非空字段名和字段值
func (model *BaseModel) GetNotNullFieldAndValues() []common.Pair[string, any] {
	return util.ListFilter(model.GetAllFieldAndValues(), func(pair common.Pair[string, any]) bool {
		return pair.Second != nil
	})
}

func (model *BaseModel) GetAllFields() []string {
	values := model.GetAllFieldAndValues()
	var result []string
	for _, value := range values {
		result = append(result, value.First)
	}
	return result
}

func (model *BaseModel) GetNotNullFields() []string {
	values := model.GetAllFieldAndValues()
	var result []string
	for _, value := range values {
		if value.Second != nil {
			result = append(result, value.First)
		}
	}
	return result
}

type JsonTime struct {
	time.Time
}