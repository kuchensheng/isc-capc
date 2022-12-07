package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/domain"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/model/api"
	"github.com/kuchensheng/capc/infrastructure/model/category"
	api2 "github.com/kuchensheng/capc/infrastructure/vo/api"
	"github.com/kuchensheng/capc/transfer/dto/api_dto"
)

var ApiDomain = func(context *gin.Context) *apiDomain {
	return &apiDomain{context, context.GetString(domain.TENANTID)}
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
	if _, ok := domain.CheckExisted(dto.Code); ok {
		return 0, common.API_CODE_EXISTS.Exception(nil)
	}
	if _, ok := domain.CheckCategory(dto.CategoryId); !ok {
		return 0, common.API_CATEGORY_IS_NULL.Exception(nil)
	}

	do := dto.Dto2DO()
	do.SetTenantId(domain.TenantId)
	return do.ID, do.Create()
}

func (domain *apiDomain) CheckExisted(code string) (api.IscCapcApiInfo, bool) {
	return api.ApiRepository.GetOne(api2.DetailVO{Code: code})
}

func (domain *apiDomain) CheckCategory(categoryId int) (category.IscCapcCategory, bool) {
	return category.CategoryRepository.GetDetail(categoryId, "")
}
