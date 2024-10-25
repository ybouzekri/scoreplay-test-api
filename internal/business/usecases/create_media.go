package usecases

import (
	"log/slog"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/repositories"
)

type CreateMediaInput struct {
	Name    string
	Picture string
	Tags    []int
}

type CreateMediaUseCase interface {
	Handle(input CreateMediaInput) (*entities.MediaEntity, error)
}

type CreateMediaInteractor struct {
	mediaRepo repositories.MediaRepository
	tagRepo   repositories.TagRepository
	logger    *slog.Logger
}

// return the interactor which handle the use case of media creation
func NewCreateMediaInteractor(mediaRepo repositories.MediaRepository, tagRepo repositories.TagRepository, logger *slog.Logger) *CreateMediaInteractor {
	return &CreateMediaInteractor{
		mediaRepo: mediaRepo,
		tagRepo:   tagRepo,
		logger:    logger,
	}
}

func (interactor *CreateMediaInteractor) Handle(input CreateMediaInput) (*entities.MediaEntity, error) {
	var tags []*entities.TagEntity
	for _, tagID := range input.Tags {
		tag, err := interactor.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	entity, err := entities.NewMediaEntity(input.Picture, input.Name, tags)
	if err != nil {
		return nil, err
	}

	return interactor.mediaRepo.Persist(entity)
}
