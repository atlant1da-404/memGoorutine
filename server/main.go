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

	mux := http.NewServeMux()
	mux.HandleFunc(utils.CreateUrl, handler.Create)
	mux.HandleFunc(utils.ListUrl, handler.List)
	mux.HandleFunc(utils.DeleteUrl, handler.Delete)
	mux.HandleFunc(utils.UpdateUrl, handler.Update)
	return http.ListenAndServe(s.addr, mux)
}
