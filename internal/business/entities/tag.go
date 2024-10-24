package entities

import "errors"

var (
	ErrEmptyTagName = errors.New("empty tag name")
)

type TagEntity struct {
	id   int
	name string
}

type TagEntityOption func(*TagEntity)

func WithTagID(id int) TagEntityOption {
	return func(entity *TagEntity) {
		entity.id = id
	}
}

func NewTagEntity(name string, options ...TagEntityOption) (*TagEntity, error) {
	if name == "" {
		return nil, ErrEmptyTagName
	}

	entity := &TagEntity{
		name: name,
	}

	for _, o := range options {
		o(entity)
	}

	return entity, nil
}

func (entity *TagEntity) ID() int {
	return entity.id
}

func (entity *TagEntity) Name() string {
	return entity.name
}
