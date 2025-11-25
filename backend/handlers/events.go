package handlers

import (
	"encoding/json"
	"kitties/orm"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// / POST /api/events
func NewEvent(w http.ResponseWriter, r *http.Request) {
	req := struct { // all fields are required
		Title       string    `json:"title"`
		Description string    `json:"description"`
		EndTime     time.Time `json:"end_time"`
		StartTime   time.Time `json:"start_time"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		// log.Println(err)
		return
	}

	event := orm.Event{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Status:      orm.EventStatusScheduled,
	}

	db := r.Context().Value(orm.DBContextKey).(*gorm.DB)
	if err := db.Save(&event).Error; err != nil {
		http.Error(w, `{"error":"couldn't save the request for some reason idk"}`, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"message": "event created",
		"event":   event,
	}) // if theres an error here u probsters cant encode the event struct

}
