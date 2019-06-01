package repository

import (
	. "gin-sample/domain"
  "errors"
)

type DesignRepository struct {
}

func NewDesignRepository() *DesignRepository {
	return new(DesignRepository)
}

func (this *DesignRepository) FindAllDesigns() ([]*Design, error) {
  var designs []*Design

  designs = append(designs, NewDesign(1, "title"))
  designs = append(designs, NewDesign(2, "title"))

	return designs, errors.New("えらーだよ＾＾")
}
