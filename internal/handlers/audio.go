package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/aneeshsunganahalli/Kumo/internal/database"
	"github.com/aneeshsunganahalli/Kumo/internal/models"
)

type AudioHandler struct {
	DB *db.DB
}

func (h *AudioHandler) CreateAudio(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var a models.Audio
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	a.UserID = userID

	if a.OriginalName == "" || a.StoredFileName == "" || a.FileType == "" {
		http.Error(w, "original_name, stored_filename, and file_type are required", http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateAudio(&a); err != nil {
		http.Error(w, "failed to create audio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

func (h *AudioHandler) GetAudio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid audio id", http.StatusBadRequest)
		return
	}

	a, err := h.DB.GetAudio(id)
	if errors.Is(err, db.ErrNotFound) {
		http.Error(w, "audio not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "failed to get audio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AudioHandler) GetAudios(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	list, err := h.DB.GetAudiosByUser(userID)
	if err != nil {
		http.Error(w, "failed to get audios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (h *AudioHandler) UpdateAudio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid audio id", http.StatusBadRequest)
		return
	}

	var a models.Audio
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.DB.UpdateAudio(id, &a); errors.Is(err, db.ErrNotFound) {
		http.Error(w, "audio not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to update audio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AudioHandler) DeleteAudio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid audio id", http.StatusBadRequest)
		return
	}

	if err := h.DB.DeleteAudio(id); errors.Is(err, db.ErrNotFound) {
		http.Error(w, "audio not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to delete audio", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
