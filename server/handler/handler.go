package handler

import (
	"net/http"
)

type Handle interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Handle
}

func NewHandler() *Handler {
	return &Handler{}
}
