package usecase

import (
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDesigns(t *testing.T) {
	uc := NewDesignUseCase(NewDesignRepository())
	var (
		total   int
		designs [10]*Design
	)
	page := 1
	err := uc.FindDesigns(&designs, &total, page)
	assert.Empty(t, err)
}

func TestFindDesign(t *testing.T) {
	uc := NewDesignUseCase(NewDesignRepository())
	page := 1
	var design Design
	designId := 1
	err := uc.FindDesign(&design, designId)
	assert.Empty(t, err)
}
