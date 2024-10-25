package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"scoreplay/internal/business/entities"
	"scoreplay/internal/business/usecases"
	"strconv"
	"strings"
)

type createMediaRequest struct {
	Name string `json:"name"`
}

type CreateMediaHandler struct {
	useCase usecases.CreateMediaUseCase
	logger  *slog.Logger
}

func NewCreateMediaHandler(useCase usecases.CreateMediaUseCase, logger *slog.Logger) *CreateMediaHandler {
	return &CreateMediaHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *CreateMediaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		response := responseError{code: http.StatusBadRequest, Message: "Error Retrieving the File"}
		response.send(w)
		return
	}
	defer file.Close()

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp("pictures", "*.png")
	if err != nil {
		response := responseError{code: http.StatusInternalServerError, Message: "Error creating image File"}
		response.send(w)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		response := responseError{code: http.StatusBadRequest, Message: "Error reading file"}
		response.send(w)
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	mediaName := r.FormValue("name")
	if mediaName == "" {
		response := responseError{code: http.StatusBadRequest, Message: "Error media name is empty"}
		response.send(w)
		return
	}

	tags := strings.Split(r.FormValue("tags"), ",")
	if len(tags) == 0 {
		response := responseError{code: http.StatusBadRequest, Message: "Error no tag specified"}
		response.send(w)
		return
	}

	var tagIDs []int
	for _, t := range tags {
		i, err := strconv.Atoi(t)
		if err != nil {
			response := responseError{code: http.StatusBadRequest, Message: "Error tag list should be a list of int"}
			response.send(w)
			return
		}
		tagIDs = append(tagIDs, i)
	}

	entity, err := h.useCase.Handle(usecases.CreateMediaInput{
		Picture: handler.Filename,
		Name:    mediaName,
		Tags:    tagIDs,
	})

	if err != nil {
		if errors.Is(err, entities.ErrEmptyMediaName) {
			response := responseError{code: http.StatusBadRequest, Message: "empty media name"}
			response.send(w)
			return
		}

		h.logger.Error("create media handler unexpected error", "error", err)
		response := responseError{code: http.StatusInternalServerError, Message: "unexpected error"}
		response.send(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMediaResponseFromEntity(entity))
}
