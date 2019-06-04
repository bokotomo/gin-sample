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

// デザイン一覧をページーションで返す
func (this *DesignRepository) FindDesigns(page int) ([]*Design, error) {
	var designs []*Design
	modelDesigns := []model.Design{}
	pageSize := 3
	offset := pageSize*page - 1

	err := DB.Offset(offset).Limit(pageSize).Find(&modelDesigns).Error
	if err != nil {
		return []*Design{}, err
	}

	for _, design := range modelDesigns {
		designs = append(designs, NewDesign(design.ID, design.Title))
	}

	return designs, nil
}
