package category

import (
	"context"
	"github.com/kuchensheng/capc/infrastructure/common"
	"github.com/kuchensheng/capc/infrastructure/model"
	"github.com/kuchensheng/capc/infrastructure/vo/category"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
)

var CategoryRepository = &categoryRepository{}

type categoryRepository struct {
	model.BaseRepository
}

func (repository *categoryRepository) GetDB(context context.Context) *gorm.DB {
	return repository.BaseRepository.GetDB(context).Table(tableName)
}

func (repository *categoryRepository) DeleteBatch(dto category.SearchVO, ctx context.Context) error {
	tx := repository.GetDB(ctx).Begin()
	defer func() {
		if x := recover(); x != nil {
			log.Warn().Msgf("发生未知异常,%v", x)
			tx.Rollback()
		}
	}()
	log.Info().Msgf("批量删除分组信息:%v", dto.Ids)
	if result := tx.Delete(NewIscCapcCategory(), dto.Ids); result.Error != nil {
		log.Warn().Msgf("无法删除分组信息，%v", result.Error)
		tx.Rollback()
		return result.Error
	} else {
		log.Info().Msgf("批量删除API信息")
	}

	result := repository.GetDB(ctx).Delete(NewIscCapcCategory(), dto.Ids)
	return result.Error
}

func (repository *categoryRepository) GetAllApp(dto category.SearchVO, context context.Context) ([]IscCapcCategory, error) {
	db := repository.GetDB(context).Where("parent_id = ?", dto.ParentId)
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

func (repository *categoryRepository) GetDetail(categoryId int, code string, context context.Context) (IscCapcCategory, bool) {
	c, f := *NewIscCapcCategory(), false
	if categoryId == 0 && code == "" {
		return c, f
	}
	db := repository.GetDB(context)
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
