package server

import (
	"encoding/json"
	"net/http"
)

const (
	start  = "start"
	stop   = "stop"
	chSize = 100
)

var handler = func(w http.ResponseWriter, r *http.Request) {

	ch := make(chan error, chSize)
	model := Storage{}
	json.NewDecoder(r.Body).Decode(&model)

	go func() {

		ch <- map[string]error{
			start: starter(model),
			stop:  stopper(model),
		}[model.Type]

	}()

	err := <-ch
	close(ch)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode("OK!")
}

var list = func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(memStorage)
}
