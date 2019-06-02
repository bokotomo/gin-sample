package usecase

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "gin-sample/usecase"
	. "gin-sample/repository"
)

func TestFindAllDesigns(t *testing.T) {
	uc := NewDesignUseCase(NewDesignRepository())
  _, err := uc.FindAllDesigns()
	assert.Empty(t, err)
}
