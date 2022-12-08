package common

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"reflect"
)

type BusinessException struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *BusinessException) Error() string {
	marshal, err := json.Marshal(e)
	if err != nil {
		log.Error().Msgf("业务异常信息序列化错误，%v", err)
		return "{}"
	}
	return string(marshal)
}

type BusinessEnum struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (be BusinessEnum) Exception(data any) *BusinessException {
	if data != nil && reflect.TypeOf(data).Kind() == reflect.TypeOf(be).Kind() {
		return data.(*BusinessException)
	}
	return &BusinessException{
		Code:    be.Code,
		Message: be.Message,
		Data:    data,
	}
}

func IsBusinessException(err error) bool {
	return nil != err && reflect.TypeOf(err).Kind() == reflect.TypeOf(&BusinessException{}).Kind()
}

func new(code int, message string) BusinessEnum {
	return BusinessEnum{Code: code, Message: message}
}

var (
	SUCCESS           = new(0, "成功")
	UNKNOWN_EXCEPTION = new(21060000, "未知异常")

	REGISTER_EXCEPTION = new(21060001, "信息注册异常")
	UPDATE_EXCEPTION   = new(21060002, "信息更新异常")
	DELETE_EXCEPTION   = new(21060003, "信息删除异常")
	ID_IS_NULL         = new(21060004, "id为空")

	API_CODE_EXISTS        = new(2106001, "api-code已存在")
	API_PATH_EXISTS        = new(2106002, "api-path已存在")
	API_NAME_EXISTS        = new(2106003, "api-name已存在")
	API_REGISTER_EXCEPTION = new(2106004, "api注册失败")
	API_NOT_EXISTS         = new(2106005, "api信息不存在")
	API_API_USED           = new(2106006, "api已被使用")
	API_CODE_IS_NULL       = new(2106007, "api code为空")
	API_NAME_IS_NULL       = new(2106008, "api name为空")
	API_PATH_IS_NULL       = new(2106008, "api path为空")
	API_CATEGORY_IS_NULL   = new(2106008, "api分组Id为空")

	NOT_ALLOW   = new(2106100, "不允许的操作")
	BAD_REQUEST = new(2106101, "参数解析异常")

	CATEGORY_CODE_EXISTS        = new(2106200, "应用标识已存在")
	CATEGORY_NAME_EXISTS        = new(2106201, "应用或分组已存在")
	CATEGORY_UNKNOWN_PARENT     = new(2106202, "无法识别的父目录")
	CATEGORY_REGISTER_EXCEPTION = new(2106203, "无法创建/更新应用或分组信息")
	CATEGORY_NOT_EXISTS         = new(2106204, "应用或分组信息不存在")
	ENV_NOT_NULL                = new(2106205, "域名信息为空")
	CATEGORY_NAME_ISNULL        = new(2106206, "应用或分组名称不能为空")
	CATEGORY_CODE_ISNULL        = new(2106207, "应用或分组标识不能为空")
	CATEGORY_ID_ISNULL          = new(2106208, "应用或分组ID不能为空")
	CATEGORY_ID_ISNOTNULL       = new(2106209, "应用或分组ID不为空")
	CATEGORY_DELETE_EXCEPTION   = new(2106210, "应用或分组删除失败")

	DYNAMIC_EXCEPTION = new(2106300, "spi引用透传失败")
	DYNAMIC_NOT_ALLOW = new(2106301, "spi引用不能通过id查询,请切换code查询")

	API_P_APIID_IS_NULL      = new(2106400, "apiId为空")
	API_P_REGISTER_EXCEPTION = new(2106401, "api参数信息注册异常")
	API_P_UPDATE_EXCEPTION   = new(2106500, "api参数信息更新异常")

	SWAGGER_PARSE_EXCEPTION = new(2106500, "swagger解析异常")
)
