package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/server"
	"github.com/rs/zerolog/log"
)

type apiHandler struct {
}

func (handler *apiHandler) RegisterHandler() {
	log.Info().Msgf("注册api管理器")
	AddHandler(&apiHandler{})
}
func (handler *apiHandler) InitView() {
	log.Info().Msgf("初始化api视图信息")
	group := server.Engine().Group("/api/orchestration/capc/api")
	{
		group.POST("", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return nil, nil
			})
		})
	}
}
