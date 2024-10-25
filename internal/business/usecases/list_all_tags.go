package usecases

import (
	"log/slog"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/repositories"
)

type ListAllTagsUseCase interface {
	Handle() ([]*entities.TagEntity, error)
}

type ListAllTagsInteractor struct {
	repository repositories.TagRepository
	logger     *slog.Logger
}

// return the interactor which handle the use case of listing all available tags
func NewListAllTagsInteractor(repository repositories.TagRepository, logger *slog.Logger) *ListAllTagsInteractor {
	return &ListAllTagsInteractor{
		repository: repository,
		logger:     logger,
	}
}

func (interactor *ListAllTagsInteractor) Handle() ([]*entities.TagEntity, error) {
	return interactor.repository.FindAll()
}
