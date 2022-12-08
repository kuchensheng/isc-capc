package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/model/api"
	"github.com/kuchensheng/capc/infrastructure/model/category"
	api2 "github.com/kuchensheng/capc/infrastructure/vo/api"
	category2 "github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/kuchensheng/capc/transfer/dto/api_dto"
	"strconv"
)

var ApiDomain = func(context *gin.Context) *apiDomain {
	return &apiDomain{context, context.GetString(common.TENANTID)}
}

type apiDomain struct {
	Context  *gin.Context
	TenantId string
}

func (domain *apiDomain) RegisterApi() (int, error) {
	op, err := domain.buildOperation()
	if err != nil {
		return 0, err
	}
	_, err = op.Create(domain.Context)
	return op.Api.ID, err
}

func (domain *apiDomain) UpdateApi() (bool, error) {
	op, err := domain.buildOperation()
	if err != nil {
		return false, err
	}
	return op.Update(domain.Context)
}

func (domain *apiDomain) DeleteApi() (bool, error) {
	strId := domain.Context.Param("id")
	code := domain.Context.Param("code")
	//路径参数不为空,因此这里不做判空处理
	var intId int
	if id, err := strconv.Atoi(strId); err != nil {
		return false, common.BAD_REQUEST.Exception(fmt.Sprintf("id = [%s] 不是int类型", strId))
	} else {
		intId = id
	}
	op := &api.ApiOperationRepository{
		Api:       api.NewIscCapcApiInfo(),
		Parameter: api.NewIscCapcApiReqResp(),
	}
	op.Api.ID = intId
	op.Api.Code = code
	op.Parameter.ApiId = intId
	op.Parameter.Code = code
	return op.Delete(domain.Context)
}

func (domain *apiDomain) buildOperation() (*api.ApiOperationRepository, error) {
	dto := &api_dto.IscApiDetailDTO{}
	if err := domain.Context.BindJSON(dto); err != nil {
		return nil, err
	}
	if _, ok := domain.CheckExisted(dto.Code, domain.Context); ok {
		return nil, common.API_CODE_EXISTS.Exception(dto.Code)
	}
	if _, ok := domain.CheckCategory(dto.CategoryId, domain.Context); !ok {
		return nil, common.CATEGORY_UNKNOWN_PARENT.Exception(nil)
	}
	if !dto.Import && dto.Type == category2.OS.GetName() {
		return nil, common.NOT_ALLOW.Exception("不允许创建OS类型的api")
	}

	apiDo, apiParamDo := dto.Dto2DO()
	apiDo.SetTenantId(domain.TenantId)
	apiParamDo.SetTenantId(domain.TenantId)
	op := &api.ApiOperationRepository{
		Api:        apiDo,
		Repository: api.ApiParameterRepository,
		Parameter:  apiParamDo,
	}
	return op, nil
}

func (domain *apiDomain) CheckExisted(code string, context context.Context) (api.IscCapcApiInfo, bool) {
	return api.ApiRepository.GetOne(api2.DetailVO{Code: code}, context)
}

func (domain *apiDomain) CheckCategory(categoryId int, context context.Context) (category.IscCapcCategory, bool) {
	return category.CategoryRepository.GetDetail(categoryId, "", context)
}
