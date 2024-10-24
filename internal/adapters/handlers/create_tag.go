package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/usecases"
)

type createTagRequest struct {
	Name string `json:"name"`
}

type CreateTagHandler struct {
	useCase usecases.CreateTagUseCase
	logger  *slog.Logger
}

func NewCreateTagHandler(useCase usecases.CreateTagUseCase, logger *slog.Logger) *CreateTagHandler {
	return &CreateTagHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *CreateTagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var requestBody createTagRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		response := responseError{code: http.StatusBadRequest, Message: "invalid request body"}
		response.send(w)
		return
	}

	entity, err := h.useCase.Handle(usecases.CreateTagInput{
		Name: requestBody.Name,
	})

	if err != nil {
		if errors.Is(err, entities.ErrEmptyTagName) {
			response := responseError{code: http.StatusBadRequest, Message: "empty tag name"}
			response.send(w)
			return
		}

		h.logger.Error("create tag handler unexpected error", "error", err)
		response := responseError{code: http.StatusInternalServerError, Message: "unexpected error"}
		response.send(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTagResponseFromEntity(entity))
}
