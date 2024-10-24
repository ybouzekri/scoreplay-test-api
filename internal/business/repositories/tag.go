package repositories

import (
	"errors"
	"scoreplay/internal/business/entities"
)

var (
	ErrTagNotFound = errors.New("tag not found")
	ErrInvalidTag  = errors.New("invalid tag")
)

type TagRepository interface {
	FindAll() ([]*entities.TagEntity, error)
	FindByID(id int) (*entities.TagEntity, error)
	Persist(tag *entities.TagEntity) (*entities.TagEntity, error)
}
