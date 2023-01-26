package usecase

import (
	"gin-sample/domain"
	"gin-sample/query/repository"
	"gin-sample/query/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDesigns(t *testing.T) {
	uc := usecase.NewDesignUseCase(repository.NewDesignRepository())
	var (
		total   int
		designs [10]domain.Design
	)
	page := 1
	err := uc.FindDesigns(&designs, &total, page)
	assert.Empty(t, err)
}

func TestFindDesign(t *testing.T) {
	uc := usecase.NewDesignUseCase(repository.NewDesignRepository())
	var design domain.Design
	designId := 1
	err := uc.FindDesign(&design, designId)
	assert.Empty(t, err)
}
