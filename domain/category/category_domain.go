package category

import (
	category2 "github.com/kuchensheng/capc/infrastructure/model/category"
	category3 "github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/kuchensheng/capc/transfer/dto/category"
	"github.com/rs/zerolog/log"
)

var CategoryDomain = &categoryDomain{}

type categoryDomain struct{}

func (domain *categoryDomain) RegisterCategory(dto *category.CategoryDTO) error {
	do := dto.Dto2Do()
	_, err := do.Create(func() (bool, error) {
		dto.ID = do.BaseModel.ID
		return true, nil
	})
	if err != nil {
		log.Warn().Msgf("分组信息注册失败,%v", err)
		return err
	}
	return nil
}

func (domain *categoryDomain) UpdateCategory(dto category.CategoryDTO) error {
	_, err := dto.Dto2Do().Update()
	if err != nil {
		log.Warn().Msgf("分组信息更新失败", err)
		return err
	}
	return nil
}

func (domain *categoryDomain) GetAllApp(dto category3.SearchVO) []category.CategoryDTO {
	app, err := category2.CategoryRepository.GetAllApp(dto)
	if err != nil {
		return nil
	}
	//todo 查询域名信息
	var result []category.CategoryDTO
	for _, model := range app {
		result = append(result, *category.Do2DTO(model))
	}
	return result
}
