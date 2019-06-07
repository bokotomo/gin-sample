package repository

import (
	. "gin-sample/domain"
	. "gin-sample/driver"
	"gin-sample/model"
)

type DesignRepository struct {
}

func NewDesignRepository() *DesignRepository {
	return new(DesignRepository)
}

/*
 * ページングトータル数を返す
 * total   ページングトータル数
 */
func (this *DesignRepository) FindDesignsTotal(total *int) error {
	return DB.Model(&model.Design{}).Count(total).Error
}

/*
 * デザイン一覧をページーションで返す
 * designs デザインドメイン一覧
 * total   ページングトータル数
 * page    ページング番号
 */
func (this *DesignRepository) FindDesigns(designs *[10]*Design, total *int, page int) error {
	resultDesigns := []model.Design{}
	if page < 0 {
		page = 1
	}
	pageSize := 10
	offset := pageSize * (page - 1)

	if err := this.FindDesignsTotal(total); err != nil {
		return err
	}

	if err := DB.Offset(offset).Limit(pageSize).Find(&resultDesigns).Error; err != nil {
		return err
	}

	for i, design := range resultDesigns {
		designs[i] = NewDesign(design.ID, design.Title)
	}

	return nil
}
