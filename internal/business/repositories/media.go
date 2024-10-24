package repositories

import (
	"errors"
	"scoreplay/internal/business/entities"
)

var (
	ErrMediaNotFound = errors.New("media not found")
)

type MediaRepository interface {
	FindByTag(tag *entities.TagEntity) (*entities.MediaEntity, error)
	Persist(media *entities.MediaEntity) (*entities.MediaEntity, error)
}
