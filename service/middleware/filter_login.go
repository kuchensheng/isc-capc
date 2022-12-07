package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/domain"
	"github.com/kuchensheng/capc/embalm"
	"github.com/kuchensheng/capc/infrastructure/common"
	"net/http"
)

func LoginFilter() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			context.JSON(http.StatusUnauthorized, common.UNKNOWN_EXCEPTION.Exception("登录校验失败,token为空"))
			context.Abort()
			return
		}
		stauts := embalm.GetUserStatus(token)
		if stauts.Data.TenantId == "" {
			context.JSON(http.StatusUnauthorized, common.UNKNOWN_EXCEPTION.Exception("租户信息校验失败"))
			context.Abort()
			return
		}
		context.Set(domain.TENANTID, stauts.Data.TenantId)
		context.Next()
	}
}
