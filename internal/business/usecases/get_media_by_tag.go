package usecases

import (
	"log/slog"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/repositories"
)

type GetMediaByTagUseCase interface {
	Handle(tagID int) ([]*entities.MediaEntity, error)
}

type GetMediaByTagInteractor struct {
	mediaRepository repositories.MediaRepository
	tagRepository   repositories.TagRepository
	logger          *slog.Logger
}

func NewGetMediaByTagInteractor(mediaRepository repositories.MediaRepository, tagRepository repositories.TagRepository, logger *slog.Logger) *GetMediaByTagInteractor {
	return &GetMediaByTagInteractor{
		mediaRepository: mediaRepository,
		tagRepository:   tagRepository,
		logger:          logger,
	}
}

func (interactor *GetMediaByTagInteractor) Handle(tagID int) ([]*entities.MediaEntity, error) {
	tag, err := interactor.tagRepository.FindByID(tagID)
	if err != nil {
		return nil, err
	}
	return interactor.mediaRepository.FindByTag(tag)
}
