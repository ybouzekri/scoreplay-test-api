package repositories

import (
	"errors"
	"log/slog"
	"scoreplay/internal/business/entities"
	"slices"
	"sync"

	"github.com/google/uuid"
)

var ErrInvalidMedia = errors.New("invalid media")

type InMemoryMediaRepository struct {
	storage []*entities.MediaEntity
	logger  *slog.Logger
	mutex   *sync.Mutex
}

func NewInMemoryMediaRepository(storage []*entities.MediaEntity, logger *slog.Logger) *InMemoryMediaRepository {
	return &InMemoryMediaRepository{
		storage: storage,
		logger:  logger,
		mutex:   &sync.Mutex{},
	}
}

func (repo *InMemoryMediaRepository) FindByTag(tag *entities.TagEntity) ([]*entities.MediaEntity, error) {
	var medias []*entities.MediaEntity
	getTagById := func(t *entities.TagEntity) bool { return t.ID() == tag.ID() }
	for _, media := range repo.storage {
		tagIndex := slices.IndexFunc(media.Tags(), getTagById)
		if tagIndex != -1 {
			medias = append(medias, media)
		}
	}
	return medias, nil
}

func (repo *InMemoryMediaRepository) Persist(media *entities.MediaEntity) (*entities.MediaEntity, error) {
	repo.logger.Debug("attempt to lock the mutex...")
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.logger.Debug("mutex locked")

	newID := uuid.New()
	newEntity, err := entities.NewMediaEntity(media.Picture(), media.Name(), media.Tags(), entities.WithMediaID(newID))
	if err != nil {
		return nil, err
	}

	repo.storage = append(repo.storage, newEntity)

	return newEntity, nil
}
