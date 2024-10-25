package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/usecases"
)

// tag response output
type tagResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newTagResponseFromEntity(entity *entities.TagEntity) tagResponse {
	return tagResponse{
		ID:   entity.ID(),
		Name: entity.Name(),
	}
}

type listTagsResponse []tagResponse

func newListTagsResponseFromEntities(entities []*entities.TagEntity) listTagsResponse {
	list := listTagsResponse{}
	for _, entity := range entities {
		list = append(list, newTagResponseFromEntity(entity))
	}

	return list
}

type ListTagsHandler struct {
	useCase usecases.ListAllTagsUseCase
	logger  *slog.Logger
}

func NewListTagsHandler(useCase usecases.ListAllTagsUseCase, logger *slog.Logger) *ListTagsHandler {
	return &ListTagsHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *ListTagsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	entities, err := h.useCase.Handle()
	if err != nil {
		h.logger.Error("list tag handler unexpected error", "error", err)
		response := responseError{code: http.StatusInternalServerError, Message: "unexpected error"}
		response.send(w)
		return
	}

	json.NewEncoder(w).Encode(newListTagsResponseFromEntities(entities))
}
