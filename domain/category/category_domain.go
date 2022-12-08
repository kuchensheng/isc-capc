package category

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kuchensheng/capc/infrastructure/common"
	category2 "github.com/kuchensheng/capc/infrastructure/model/category"
	category3 "github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/kuchensheng/capc/transfer/dto/category"
	"github.com/kuchensheng/capc/util"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

var CategoryDomain = func(context *gin.Context) *categoryDomain {
	return &categoryDomain{context, context.GetString(common.TENANTID)}
}

type categoryDomain struct {
	Context  *gin.Context
	TenantId string
}

func (domain *categoryDomain) RegisterCategory() (int, error) {
	dto := getRawData2DTO(domain.Context)
	do := dto.Dto2Do()
	if _, err := do.Create(domain.Context, nil); err != nil {
		log.Warn().Msgf("分组信息注册失败,%v", err)
		return 0, err
	}
	dto.ID = do.ID
	return do.ID, nil
}

func (domain *categoryDomain) UpdateCategory(dto category.CategoryDTO) error {
	_, err := dto.Dto2Do().Update(domain.Context, nil)
	if err != nil {
		log.Warn().Msgf("分组信息更新失败", err)
		return err
	}
	return nil
}

func (domain *categoryDomain) GetAllApp() ([]category.CategoryDTO, error) {
	app, err := category2.CategoryRepository.GetAllApp(getSearchDTO(domain.Context), domain.Context)
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

func (domain *categoryDomain) DeleteCategory() (bool, error) {

	if dto, err := domain.GetDetailById(); err != nil {
		return false, err
	} else {
		children := domain.getChildren(dto.ID)
		ids := util.Map(children, func(t category2.IscCapcCategory) int { return t.ID })
		ids = append(ids, dto.ID)
		//批量删除
		if err = category2.CategoryRepository.DeleteBatch(category3.SearchVO{Ids: ids}, domain.Context); err != nil {
			return false, common.BAD_REQUEST.Exception("批量删除失败")
		}
		return true, nil
	}
}

//getChildren 获取子孙节点，这里采用递归查询数据库的做法，原因：分组的量很少，层级也不会很深
func (domain *categoryDomain) getChildren(id int) []category2.IscCapcCategory {
	var result []category2.IscCapcCategory
	if app, err := category2.CategoryRepository.GetAllApp(category3.SearchVO{ParentId: id}, domain.Context); err != nil || len(app) == 0 {
		return result
	} else {
		result = append(result, app...)
		for _, capcCategory := range app {
			result = append(result, domain.getChildren(capcCategory.ID)...)
		}
	}
	return result
}

func (domain *categoryDomain) GetDetailById() (*category.CategoryDTO, error) {
	strCategoryId := domain.Context.Param("id")
	var categoryId int
	if id, err := strconv.Atoi(strCategoryId); err != nil {
		return nil, common.BAD_REQUEST.Exception(fmt.Sprintf("id不是整形，id = %s", strCategoryId))
	} else {
		categoryId = id
	}
	code := domain.Context.Param("code")
	if model, ok := category2.CategoryRepository.GetDetail(categoryId, code, domain.Context); !ok {
		return nil, common.CATEGORY_NOT_EXISTS.Exception(nil)
	} else {
		return category.Do2DTO(model), nil
	}
}

func getSearchDTO(ctx *gin.Context) category3.SearchVO {
	search := &category3.SearchVO{}
	ctx.ShouldBind(search)

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
