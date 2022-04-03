package server

import (
	"atlant1da-404/server/handler"
	"atlant1da-404/server/utils"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer() *Server {
	return &Server{
		addr: utils.Hardcode,
	}
}

func (s Server) Run() error {

	h := handler.NewHandler()

	mux := http.NewServeMux()
	mux.HandleFunc(utils.CreateUrl, h.Create)
	mux.HandleFunc(utils.ListUrl, h.List)
	mux.HandleFunc(utils.DeleteUrl, h.Delete)
	mux.HandleFunc(utils.UpdateUrl, h.Update)
	return http.ListenAndServe(s.addr, mux)
}
