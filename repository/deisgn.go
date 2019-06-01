package repository

import (
	. "gin-sample/domain"
)

type DesignRepository struct {
}

func NewDesignRepository() *DesignRepository {
	return new(DesignRepository)
}

func (this *DesignRepository) FindAllDesigns() ([]*Design, error) {
  var designs []*Design

  design := NewDesign(1, "OK")
  designs = append(designs, design)
  designs = append(designs, design)
	return designs, nil
}
