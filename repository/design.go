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
func (this *DesignRepository) FindDesigns(designs *[10]Design, total *int, page int) error {
	if page < 0 {
		page = 1
	}
	pageSize := 10
	offset := pageSize * (page - 1)

	if err := this.FindDesignsTotal(total); err != nil {
		return err
	}

	ds := []model.Design{}
	if err := DB.Offset(offset).Limit(pageSize).Find(&ds).Error; err != nil {
		return err
	}

	for i, d := range ds {
		designs[i].SetPickup(d.ID, d.Title, d.Good, d.Thumbnail)
	}

	return nil
}

/*
 * デザイン詳細を返す
 * designs  デザインドメイン一覧
 * designId デザインID
 */
func (this *DesignRepository) FindDesign(design *Design, designId int) error {
	d := model.Design{}
	if err := DB.First(&d, designId).Error; err != nil {
		return err
	}
	design.Set(d.ID, d.Title, d.Text, d.Good, d.Comments, d.CreatedAt.Format("2006-01-02 15:4"), d.Thumbnail)

	return nil
}
