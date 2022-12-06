package category

import (
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
)

var CategoryRepository *categoryRepository = &categoryRepository{model.BaseRepository{DB: connetor.Db, TableName: tableName}}

type categoryRepository struct {
	model.BaseRepository
}

func (repository *categoryRepository) GetDB() *gorm.DB {
	return repository.BaseRepository.GetDB()
}

func (repository *categoryRepository) GetAllApp(dto category.SearchVO) ([]CategoryModel, error) {
	db := repository.GetDB().Where("parent_id = ?", dto.ParentId)
	if dto.Name != "" && strings.Trim(dto.Name, " ") != "" {
		db = db.Where("name LIKE ?", "%"+dto.Name+"%")
	}
	if dto.Type != 0 {
		db = db.Where("type = ?", dto.Type)
	}
	if dto.Ids != nil {
		db = db.Where("id IN ?", dto.Ids)
	}
	if dto.Codes != nil {
		db = db.Where("code IN ?", dto.Codes)
	}

	var result []CategoryModel
	rows, err := db.Table(repository.TableName).Rows()
	defer rows.Close()
	if err != nil {
		log.Warn().Msgf("分组信息查询失败,%v", err)
		return nil, common.BAD_REQUEST.Exception(err.Error())
	}

	for rows.Next() {
		var item CategoryModel
		if err := db.ScanRows(rows, &item); err != nil {
			log.Warn().Msgf("信息扫描失败,%v", err)
			continue
		} else {
			result = append(result, item)
		}
	}
	return result, nil

}
