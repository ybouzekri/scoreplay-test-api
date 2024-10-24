package entities_test

import (
	"scoreplay/internal/business/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTagEntityWithOptions(t *testing.T) {
	id := 42
	name := "tag name"

	entity, err := entities.NewTagEntity(
		name,
		entities.WithTagID(id),
	)

	assert.NoError(t, err)
	assert.Equal(t, id, entity.ID())
	assert.Equal(t, name, entity.Name())
}

func TestNewTagEntityShouldNotAcceptEmptyTagName(t *testing.T) {
	_, err := entities.NewTagEntity("")
	assert.ErrorIs(t, err, entities.ErrEmptyTagName)
}
