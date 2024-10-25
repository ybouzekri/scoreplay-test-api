package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/usecases"
	"strconv"
)

// media response output
type mediaResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	FileUrl string   `json:"fileUrl"`
}

func newMediaResponseFromEntity(entity *entities.MediaEntity) mediaResponse {
	var tagNames []string
	for _, tag := range entity.Tags() {
		tagNames = append(tagNames, tag.Name()) // note the = instead of :=
	}

	return mediaResponse{
		ID:      entity.ID().String(),
		Name:    entity.Name(),
		Tags:    tagNames,
		FileUrl: fmt.Sprintf("http://0.0.0.0:8888/images/%s", entity.Picture()),
	}
}

type getMediaByTagResponse []mediaResponse

func newGetMediaByTagResponseFromEntities(entities []*entities.MediaEntity) getMediaByTagResponse {
	list := getMediaByTagResponse{}
	for _, entity := range entities {
		list = append(list, newMediaResponseFromEntity(entity))
	}

	return list
}

type GetMediaByTagHandler struct {
	useCase usecases.GetMediaByTagUseCase
	logger  *slog.Logger
}

func NewGetMediaByTagHandler(useCase usecases.GetMediaByTagUseCase, logger *slog.Logger) *GetMediaByTagHandler {
	return &GetMediaByTagHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *GetMediaByTagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// Get tag query param
	tag := r.URL.Query().Get("tag")
	if tag == "" {
		response := responseError{code: http.StatusBadRequest, Message: "Error tag query should be provided"}
		response.send(w)
		return
	}
	// validate its an int
	tagID, err := strconv.Atoi(tag)
	if err != nil {
		response := responseError{code: http.StatusBadRequest, Message: "Error tag should be an integer"}
		response.send(w)
		return
	}
	entities, err := h.useCase.Handle(tagID)
	if err != nil {
		response := responseError{code: http.StatusInternalServerError, Message: "unexpected error"}
		response.send(w)
		return
	}

	json.NewEncoder(w).Encode(newGetMediaByTagResponseFromEntities(entities))
}
