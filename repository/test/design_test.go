package repository

import (
	. "gin-sample/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllDesigns(t *testing.T) {
	repository := NewDesignRepository()
	_, err := repository.FindAllDesigns()
	assert.Empty(t, err)
}
