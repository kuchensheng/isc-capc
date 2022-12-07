package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/infrastructure/common"
	"net/http"
)

var AllHandler []InitHandler

type InitHandler interface {
	//InitView 初始化视图
	InitView()
}

//RegisterHandler 注册处理器
func RegisterHandler(handler InitHandler) {
	AllHandler = append(AllHandler, handler)
}

//InitAllView 初始化所有的视图信息
func InitAllView() {
	for _, handler := range AllHandler {
		handler.InitView()
	}
}

func handlerBusiness(context *gin.Context, handler func(context *gin.Context) (any, error)) {
	if result, err := handler(context); err != nil {
		context.JSON(http.StatusBadRequest, common.BAD_REQUEST.Exception(err))
		context.Abort()
	} else {
		context.JSON(http.StatusOK, common.SUCCESS.Exception(result))
		return
	}
}

func getTenantId(context *gin.Context) (string, error) {
	header := context.GetHeader("token")
	if header == "" {
		context.JSON(http.StatusUnauthorized, common.UNKNOWN_EXCEPTION.Exception("登录校验失败"))
		context.Abort()
	}
	return header, nil
}
