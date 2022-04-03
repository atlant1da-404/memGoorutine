package handler

import (
	"atlant1da-404/server/service"
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"encoding/json"
	"errors"
	"net/http"
)

var Create = func(w http.ResponseWriter, r *http.Request) {

	ch := make(chan error, utils.ChSize)
	model := storage.Storage{}
	json.NewDecoder(r.Body).Decode(&model)

	go func() {

		ch <- map[string]error{
			utils.Start: service.Add(model),
			utils.Stop:  service.Stop(model),
		}[model.Type]

	}()

	err := <-ch
	close(ch)
	if err != nil {
		utils.HttpError(w, err)
		return
	}

	json.NewEncoder(w).Encode(utils.OK)
}

var List = func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.MemStorage)
}

var Delete = func(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	requestId := queryParams.Get(utils.ReqId)

	model := storage.Storage{RequestId: requestId}

	if !service.CheckInStorage(model) {
		utils.HttpError(w, errors.New(utils.ErrNotFound))
		return
	}

	delete(storage.MemStorage, requestId)
	json.NewEncoder(w).Encode(utils.OK)
}

var Update = func(w http.ResponseWriter, r *http.Request) {

	model := storage.Storage{}
	json.NewDecoder(r.Body).Decode(&model)

	if !service.CheckInStorage(model) {
		utils.HttpError(w, errors.New(utils.ErrNotFound))
		return
	}

	service.Update(model)
	json.NewEncoder(w).Encode(utils.OK)
}
