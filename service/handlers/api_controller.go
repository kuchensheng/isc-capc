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
	group := server.Engine().Group("/api/orchestration/capc/api")
	group.Use(middleware.LoginFilter())
	{
		//创建API信息
		group.POST("", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return api.ApiDomain(context).RegisterApi()
			})
		})
	}
}
