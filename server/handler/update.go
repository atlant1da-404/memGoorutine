package handler

import (
	"atlant1da-404/server/service"
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {

	model := storage.Storage{}
	json.NewDecoder(r.Body).Decode(&model)
	s := service.NewService(model)

	if !s.CheckInStorage() {
		utils.HttpError(w, errors.New(utils.ErrNotFound))
		return
	}

	s.Update()
	json.NewEncoder(w).Encode(utils.OK)
}
