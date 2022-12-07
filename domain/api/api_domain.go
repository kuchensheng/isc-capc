package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/domain"
	"github.com/kuchensheng/capc/transfer/dto/api_dto"
)

var ApiDomain = func(context *gin.Context) *apiDomain {
	return &apiDomain{context, context.GetHeader(domain.TENANTID)}
}

type apiDomain struct {
	Context  *gin.Context
	TenantId string
}

func (domain *apiDomain) RegisterApi() (int, error) {
	dto := &api_dto.IscApiInfoDTO{}
	if err := domain.Context.BindJSON(dto); err != nil {
		return 0, err
	}
	do := dto.Dto2DO()
	do.SetTenantId(domain.TenantId)
	return do.ID, do.Create()
}
