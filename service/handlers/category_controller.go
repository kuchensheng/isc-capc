package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/server"
	"github.com/kuchensheng/capc/domain/category"
	"github.com/kuchensheng/capc/infrastructure/common"
	category3 "github.com/kuchensheng/capc/infrastructure/vo/category"
	category2 "github.com/kuchensheng/capc/transfer/dto/category"
	"github.com/kuchensheng/capc/util"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type categoryHandler struct{}

func (handler *categoryHandler) RegisterHandler() {
	log.Info().Msgf("分组控制器注册...")
	AddHandler(&categoryHandler{})
}

func (handler *categoryHandler) InitView() {
	log.Info().Msgf("初始化分组控制器")
	g := server.Engine().Group("/api/orchestration/capc/category")
	{
		//注册应用或者分组信息
		g.POST("", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				dto := getRawData2DTO(context)
				return dto.ID, category.CategoryDomain.RegisterCategory(&dto)
			})
		})
		//获取应用或分组列表
		g.GET("/category/list", func(context *gin.Context) {
			handlerBusiness(context, func(context *gin.Context) (any, error) {
				return category.CategoryDomain.GetAllApp(getSearchDTO(context)), nil
			})
		})
	}
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

func getRawData2DTO(ctx *gin.Context) (dto category2.CategoryDTO) {
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
