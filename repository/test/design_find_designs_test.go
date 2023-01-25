package repository

import (
	"gin-sample/domain"
	"gin-sample/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllDesigns(t *testing.T) {
	repo := repository.NewDesignRepository()
	page := 1
	var (
		total   int
		designs [10]domain.Design
	)
	err := repo.FindDesigns(&designs, &total, page)
	assert.Empty(t, err)
}
