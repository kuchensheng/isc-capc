package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/infrastructure/common"
	"net/http"
)

var AllHandler []InitHandler

type InitHandler interface {
	//RegisterHandler 注册处理器
	RegisterHandler()
	//InitView 初始化视图
	InitView()
}

//AddHandler 注册处理器
func AddHandler(handler InitHandler) {
	AllHandler = append(AllHandler, handler)
}

//InitAllView 初始化所有的视图信息
func InitAllView() {
	categoryHandler := &categoryHandler{}
	categoryHandler.InitView()

	apiHandler := &apiHandler{}
	apiHandler.InitView()
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
