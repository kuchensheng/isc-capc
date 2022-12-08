package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/server"
	"github.com/kuchensheng/capc/domain/api"
	"github.com/kuchensheng/capc/service/middleware"
	"github.com/rs/zerolog/log"
)

type apiHandler struct {
}

func init() {
	RegisterHandler(&apiHandler{})
}

func (handler *apiHandler) InitView() {
	log.Info().Msgf("初始化api视图信息")
	group := server.Engine().Group("/api/orchestration/capc")
	group.Use(middleware.LoginFilter())
	{
		//创建API信息
		group.POST("/api", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return api.ApiDomain(context).RegisterApi()
			})
		})

		//更新API信息
		group.PUT("/api", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return api.ApiDomain(context).UpdateApi()
			})
		})
		//删除API信息
		group.DELETE("/api/:id", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return api.ApiDomain(context).DeleteApi()
			})
		})
		//删除API信息
		group.DELETE("/api/code/:code", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return api.ApiDomain(context).DeleteApi()
			})
		})
	}
}
