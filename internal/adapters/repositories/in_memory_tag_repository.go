package repositories

import (
	"log/slog"
	"scoreplay/internal/business/entities"
	"sync"
)

type InMemoryTagRepository struct {
	storage []*entities.TagEntity
	logger  *slog.Logger
	mutex   *sync.Mutex
}

func NewInMemoryTagRepository(storage []*entities.TagEntity, logger *slog.Logger) *InMemoryTagRepository {
	return &InMemoryTagRepository{
		storage: storage,
		logger:  logger,
		mutex:   &sync.Mutex{},
	}
}

func (repo *InMemoryTagRepository) FindAll() ([]*entities.TagEntity, error) {
	return repo.storage, nil
}

func (repo *InMemoryTagRepository) FindByID(id int) (*entities.TagEntity, error) {
	panic("not implemented")
}

func (repo *InMemoryTagRepository) Persist(tag *entities.TagEntity) (*entities.TagEntity, error) {
	repo.logger.Debug("attempt to lock the mutex...")
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.logger.Debug("mutex locked")

	var last int
	for _, entity := range repo.storage {
		if entity.ID() > last {
			last = entity.ID()
		}
	}

	newID := last + 1
	newEntity, err := entities.NewTagEntity(tag.Name(), entities.WithTagID(newID))
	if err != nil {
		return nil, err
	}

	repo.storage = append(repo.storage, newEntity)

	return newEntity, nil
}
