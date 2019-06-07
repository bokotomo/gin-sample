package repository

import (
	. "gin-sample/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllDesigns(t *testing.T) {
	repository := NewDesignRepository()
	page := 1
	_, _, err := repository.FindDesigns(page)
	assert.Empty(t, err)
}
