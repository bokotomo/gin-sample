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

func (this *DesignRepository) FindAllDesigns() ([]*Design, error) {
  var designs []*Design
  modelDesigns := []model.Design{}

  if err := DB.Find(&modelDesigns).Error; err != nil {
    return []*Design{}, err
  }

  for _, design := range modelDesigns {
    designs = append(designs, NewDesign(design.ID, design.Title))
  }

	return designs, nil
}
