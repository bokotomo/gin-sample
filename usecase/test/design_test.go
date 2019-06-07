package usecase

import (
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllDesigns(t *testing.T) {
	uc := NewDesignUseCase(NewDesignRepository())
	page := 1
	_, _, err := uc.FindDesigns(page)
	assert.Empty(t, err)
}
