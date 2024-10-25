package usecases

import (
	"log/slog"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/repositories"
)

type CreateTagInput struct {
	Name string
}

type CreateTagUseCase interface {
	Handle(input CreateTagInput) (*entities.TagEntity, error)
}

type CreateTagInteractor struct {
	repo   repositories.TagRepository
	logger *slog.Logger
}

// return the interactor which handle the use case of media creation
func NewCreateTagInteractor(repo repositories.TagRepository, logger *slog.Logger) *CreateTagInteractor {
	return &CreateTagInteractor{
		repo:   repo,
		logger: logger,
	}
}

func (interactor *CreateTagInteractor) Handle(input CreateTagInput) (*entities.TagEntity, error) {
	entity, err := entities.NewTagEntity(input.Name)
	if err != nil {
		return nil, err
	}

	return interactor.repo.Persist(entity)
}
