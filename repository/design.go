package repository

import (
	"gin-sample/domain"
	"gin-sample/driver"
	"gin-sample/model"
)

// DesignRepository is
type DesignRepository struct{}

// NewDesignRepository is
func NewDesignRepository() *DesignRepository {
	return new(DesignRepository)
}

// FindDesignsTotal ページングトータル数を返す
func (dr *DesignRepository) FindDesignsTotal(total *int) error {
	return driver.DB.Model(&model.Design{}).Count(total).Error
}

// FindDesigns デザイン一覧をページーションで返す
func (dr *DesignRepository) FindDesigns(designs *[10]domain.Design, total *int, page int) error {
	if page < 0 {
		page = 1
	}
	pageSize := 10
	offset := pageSize * (page - 1)

	if err := dr.FindDesignsTotal(total); err != nil {
		return err
	}

	ds := []model.Design{}
	if err := driver.DB.Offset(offset).Limit(pageSize).Find(&ds).Error; err != nil {
		return err
	}

	for i, d := range ds {
		designs[i].SetPickup(d.ID, d.Title, d.Good, d.Thumbnail)
	}

	return nil
}

// FindDesign デザイン詳細を返す
func (dr *DesignRepository) FindDesign(design *domain.Design, designId int) error {
	d := model.Design{}
	if err := driver.DB.First(&d, designId).Error; err != nil {
		return err
	}
	design.Set(d.ID, d.Title, d.Text, d.Good, d.Comments, d.CreatedAt.Format("2006-01-02 15:4"), d.Thumbnail)

	return nil
}
