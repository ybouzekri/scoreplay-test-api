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
	mediaRepository := repositories.NewInMemoryMediaRepository([]*entities.MediaEntity{}, logger)

	listAllTagsUseCase := usecases.NewListAllTagsInteractor(tagRepository, logger)
	createTagUseCase := usecases.NewCreateTagInteractor(tagRepository, logger)

	getMediaByTagUseCase := usecases.NewGetMediaByTagInteractor(mediaRepository, tagRepository, logger)

	listTagsHandler := handlers.NewListTagsHandler(listAllTagsUseCase, logger)
	createTagHandler := handlers.NewCreateTagHandler(createTagUseCase, logger)

	getMediaByTagHandler := handlers.NewGetMediaByTagHandler(getMediaByTagUseCase, logger)

	router.Handle("GET /tags", listTagsHandler)
	router.Handle("POST /tags", createTagHandler)

	router.Handle("GET /media", getMediaByTagHandler)

	return router
}
