package server

import (
	"net/http"
)

type Server struct {
	addr string
}

const (
	hardcode  = "localhost:8000"
	createUrl = "/"
	listUrl   = "/l"
)

func NewServer() *Server {
	return &Server{
		addr: hardcode,
	}
}

func (s Server) Run() error {

	mux := http.NewServeMux()
	mux.HandleFunc(createUrl, handler)
	mux.HandleFunc(listUrl, list)
	return http.ListenAndServe(s.addr, mux)
}
