package category

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/domain"
	"github.com/kuchensheng/capc/infrastructure/common"
	category2 "github.com/kuchensheng/capc/infrastructure/model/category"
	category3 "github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/kuchensheng/capc/transfer/dto/category"
	"github.com/kuchensheng/capc/util"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

var CategoryDomain = func(context *gin.Context) *categoryDomain {
	return &categoryDomain{context, context.GetHeader(domain.TENANTID)}
}

type categoryDomain struct {
	Context  *gin.Context
	TenantId string
}

func (domain *categoryDomain) RegisterCategory() (int, error) {
	dto := getRawData2DTO(domain.Context)
	do := dto.Dto2Do()
	_, err := do.Create(func() (bool, error) {
		dto.ID = do.BaseModel.ID
		do.SetTenantId(domain.TenantId)
		return true, nil
	})
	if err != nil {
		log.Warn().Msgf("分组信息注册失败,%v", err)
		return 0, err
	}
	return do.ID, nil
}

func (domain *categoryDomain) UpdateCategory(dto category.CategoryDTO) error {
	_, err := dto.Dto2Do().Update()
	if err != nil {
		log.Warn().Msgf("分组信息更新失败", err)
		return err
	}
	return nil
}

func (domain *categoryDomain) GetAllApp() ([]category.CategoryDTO, error) {
	app, err := category2.CategoryRepository.GetAllApp(getSearchDTO(domain.Context))
	if err != nil {
		return nil, err
	}
	//todo 查询域名信息
	var result []category.CategoryDTO
	for _, model := range app {
		result = append(result, *category.Do2DTO(model))
	}
	return result, nil
}

func getSearchDTO(ctx *gin.Context) category3.SearchVO {
	search := &category3.SearchVO{}
	if v, ok := ctx.GetQuery("name"); ok {
		search.Name = v
	}

	if v, ok := ctx.GetQuery("type"); ok {
		search.Type = category3.CategoryType(util.StrToInt(v))
	}
	if v, ok := ctx.GetQuery("parentId"); ok {
		search.ParentId = util.StrToInt(v)
	}

	if v, ok := ctx.GetQuery("codes"); ok {
		v = strings.ReplaceAll(strings.ReplaceAll(v, "[", ""), "]", "")
		search.Codes = strings.Split(v, ",")
	}
	if v, ok := ctx.GetQuery("ids"); ok {
		search.Ids = func(strArray string) []int {
			var result []int
			strArray = strings.ReplaceAll(strings.ReplaceAll(strArray, "[", ""), "]", "")
			for _, s := range strings.Split(strArray, ",") {
				result = append(result, util.StrToInt(s))
			}
			return result
		}(v)
	}
	return *search
}

func getRawData2DTO(ctx *gin.Context) (dto category.CategoryDTO) {
	data, err := ctx.GetRawData()
	if err != nil {
		log.Warn().Msgf("请求体读取异常,%v", err)
		ctx.JSON(http.StatusBadRequest, common.BAD_REQUEST.Exception(err.Error()))
		return
	} else if err = json.Unmarshal(data, &dto); err != nil {
		log.Warn().Msgf("请求体解析异常,%v", err)
		ctx.JSON(http.StatusBadRequest, common.BAD_REQUEST.Exception(err.Error()))
		return
	}
	return
}
