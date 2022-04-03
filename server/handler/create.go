package handler

import (
	"atlant1da-404/server/service"
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"encoding/json"
	"net/http"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

	ch := make(chan error, utils.ChSize)

	model := storage.Storage{}
	json.NewDecoder(r.Body).Decode(&model)
	s := service.NewService(model)

	go func() {

		ch <- map[string]error{
			utils.Start: s.Add(),
			utils.Stop:  s.Stop(),
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
