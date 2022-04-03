package handler

import (
	"atlant1da-404/server/service"
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	requestId := queryParams.Get(utils.ReqId)
	model := storage.Storage{RequestId: requestId}
	s := service.NewService(model)

	if !s.CheckInStorage() {
		utils.HttpError(w, errors.New(utils.ErrNotFound))
		return
	}

	delete(storage.MemStorage, requestId)
	json.NewEncoder(w).Encode(utils.OK)
}
