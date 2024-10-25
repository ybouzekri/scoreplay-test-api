package entities

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyPicture   = errors.New("empty picture")
	ErrEmptyMediaName = errors.New("empty media name")
	ErrEmptyTagList   = errors.New("empty tag list")
	ErrInvalidTag     = errors.New("invalid tag")
)

type MediaEntity struct {
	id      uuid.UUID
	picture string
	name    string
	tags    []*TagEntity
}

type MediaEntityOption func(*MediaEntity)

func WithMediaID(id uuid.UUID) MediaEntityOption {
	return func(m *MediaEntity) {
		m.id = id
	}
}

// creates a new media entity, the optional functions are here to ease the construction of entities that are valide but in incomplete states
// such as during creation the entity ID is determined by the database, so we should be able to create the entity without the ID
// However the ID is the entity's identity and must be immutable, the option pattern in Go solves this issue by making certain parameters optional
func NewMediaEntity(
	picture string,
	name string,
	tags []*TagEntity,
	options ...MediaEntityOption,
) (*MediaEntity, error) {
	if len(picture) == 0 {
		return nil, ErrEmptyPicture
	}

	if name == "" {
		return nil, ErrEmptyMediaName
	}

	if len(tags) == 0 {
		return nil, ErrEmptyTagList
	}

	for _, t := range tags {
		if t.id == 0 || t.name == "" {
			return nil, ErrInvalidTag
		}
	}

	media := &MediaEntity{
		picture: picture,
		name:    name,
		tags:    tags,
	}

	for _, o := range options {
		o(media)
	}

	return media, nil
}

func (entity *MediaEntity) ID() uuid.UUID {
	return entity.id
}

func (entity *MediaEntity) Picture() string {
	return entity.picture
}

func (entity *MediaEntity) Name() string {
	return entity.name
}

func (entity *MediaEntity) Tags() []*TagEntity {
	return entity.tags
}
