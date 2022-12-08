package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/server"
	"github.com/kuchensheng/capc/domain/category"
	"github.com/kuchensheng/capc/service/middleware"
	"github.com/rs/zerolog/log"
)

type categoryHandler struct{}

func init() {
	RegisterHandler(&categoryHandler{})
}

func (handler *categoryHandler) InitView() {
	log.Info().Msgf("初始化分组控制器")
	g := server.Engine().Group("/api/orchestration/capc/category")
	g.Use(middleware.LoginFilter())
	{
		//注册应用或者分组信息
		g.POST("", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return category.CategoryDomain(context).RegisterCategory()
			})
		})
		//获取应用或分组列表
		g.GET("/list", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return category.CategoryDomain(context).GetAllApp()
			})
		})
		//删除应用或分组
		g.DELETE("/del/:id", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return category.CategoryDomain(context).DeleteCategory()
			})
		})
	}
}
