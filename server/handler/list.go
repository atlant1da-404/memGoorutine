package handler

import (
	"atlant1da-404/server/storage"
	"encoding/json"
	"net/http"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.MemStorage)
}
