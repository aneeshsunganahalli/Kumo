package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/aneeshsunganahalli/Kumo/internal/database"
	"github.com/aneeshsunganahalli/Kumo/internal/models"
)

type PreferenceHandler struct {
	DB *db.DB
}

func (h *PreferenceHandler) CreatePreferences(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var pref models.UIPreference
	if err := json.NewDecoder(r.Body).Decode(&pref); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	pref.UserID = userID

	if err := h.DB.CreatePreferences(&pref); err != nil {
		http.Error(w, "failed to create preferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pref)
}

func (h *PreferenceHandler) GetPreferences(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	pref, err := h.DB.GetPreferences(userID)
	if errors.Is(err, db.ErrNotFound) {
		http.Error(w, "preferences not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "failed to get preferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pref)
}

func (h *PreferenceHandler) UpdatePreferences(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var pref models.UIPreference
	if err := json.NewDecoder(r.Body).Decode(&pref); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.DB.UpdatePreferences(userID, &pref); errors.Is(err, db.ErrNotFound) {
		http.Error(w, "preferences not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to update preferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pref)
}
