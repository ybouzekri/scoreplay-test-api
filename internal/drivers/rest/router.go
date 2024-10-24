package rest

import (
	"log/slog"
	"net/http"
	"scoreplay/internal/adapters/handlers"
	"scoreplay/internal/adapters/repositories"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/usecases"
)

func NewRouter(logger *slog.Logger) *http.ServeMux {
	router := http.NewServeMux()

	tagRepository := repositories.NewInMemoryTagRepository([]*entities.TagEntity{}, logger)

	listAllTagsUseCase := usecases.NewListAllTagsInteractor(tagRepository, logger)
	createTagUseCase := usecases.NewCreateTagInteractor(tagRepository, logger)

	listTagsHandler := handlers.NewListTagsHandler(listAllTagsUseCase, logger)
	createTagHandler := handlers.NewCreateTagHandler(createTagUseCase, logger)

	router.Handle("GET /tags", listTagsHandler)
	router.Handle("POST /tags", createTagHandler)

	return router
}
