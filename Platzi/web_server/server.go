package main

import (
	"net/http"
)

type Server struct{
	port string
	router *Router
}

//funcion a ser utilizada de manera publica por la mayuscula, en el mismo paquete
func NewServer(port string) *Server{
	return &Server{
		port: port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc){
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)
	if error != nil{
		return nil
	}
	return nil
}