package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindAllDesigns(t *testing.T) {
	repository := NewDesignRepository()
  _, err := repository.FindAllDesigns()
	assert.Empty(t, err)
}
