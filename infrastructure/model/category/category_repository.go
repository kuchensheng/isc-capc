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

var CategoryRepository = &categoryRepository{model.BaseRepository{DB: connetor.Db.Table(tableName)}}

type categoryRepository struct {
	model.BaseRepository
}

func (repository *categoryRepository) GetDB() *gorm.DB {
	return repository.BaseRepository.GetDB()
}

func (repository *categoryRepository) GetAllApp(dto category.SearchVO) ([]IscCapcCategory, error) {
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

	var result []IscCapcCategory
	rows, err := db.Rows()
	defer rows.Close()
	if err != nil {
		log.Warn().Msgf("分组信息查询失败,%v", err)
		return nil, common.BAD_REQUEST.Exception(err.Error())
	}

	for rows.Next() {
		var item IscCapcCategory
		if err := db.ScanRows(rows, &item); err != nil {
			log.Warn().Msgf("信息扫描失败,%v", err)
			continue
		} else {
			result = append(result, item)
		}
	}
	return result, nil

}

func (repository *categoryRepository) GetDetail(categoryId int, code string) (IscCapcCategory, bool) {
	c, f := IscCapcCategory{}, false
	if categoryId == 0 && code == "" {
		return c, f
	}
	db := repository.GetDB()
	if categoryId != 0 {
		db = db.Where("id = ?", categoryId)
	}
	if code != "" {
		db = db.Where("code = ?", code)
	}
	db = db.Take(&c)
	if db.Error != nil || db.RowsAffected < 1 {
		log.Warn().Msgf("未获取到分组信息,error=%v,count=%d", db.Error, db.RowsAffected)
		return c, f
	}
	return c, true
}
