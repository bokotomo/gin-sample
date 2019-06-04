package usecase

import (
	. "gin-sample/repository"
	. "gin-sample/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllDesigns(t *testing.T) {
	uc := NewDesignUseCase(NewDesignRepository())
	_, err := uc.FindAllDesigns()
	assert.Empty(t, err)
}
