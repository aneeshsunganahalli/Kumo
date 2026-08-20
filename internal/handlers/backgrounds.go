package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/aneeshsunganahalli/Kumo/internal/database"
	"github.com/aneeshsunganahalli/Kumo/internal/models"
)

type BackgroundHandler struct {
	DB *db.DB
}

func (h *BackgroundHandler) CreateBackground(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var b models.Background
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	b.UserID = userID

	if b.OriginalName == "" || b.StoredFileName == "" || b.FileType == "" {
		http.Error(w, "original_name, stored_filename, and file_type are required", http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateBackground(&b); err != nil {
		http.Error(w, "failed to create background", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func (h *BackgroundHandler) GetBackground(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid background id", http.StatusBadRequest)
		return
	}

	b, err := h.DB.GetBackground(id)
	if errors.Is(err, db.ErrNotFound) {
		http.Error(w, "background not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "failed to get background", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func (h *BackgroundHandler) GetBackgrounds(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	list, err := h.DB.GetBackgroundsByUser(userID)
	if err != nil {
		http.Error(w, "failed to get backgrounds", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (h *BackgroundHandler) UpdateBackground(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid background id", http.StatusBadRequest)
		return
	}

	var b models.Background
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.DB.UpdateBackground(id, &b); errors.Is(err, db.ErrNotFound) {
		http.Error(w, "background not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to update background", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func (h *BackgroundHandler) DeleteBackground(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid background id", http.StatusBadRequest)
		return
	}

	if err := h.DB.DeleteBackground(id); errors.Is(err, db.ErrNotFound) {
		http.Error(w, "background not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to delete background", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
