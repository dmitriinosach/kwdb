package http

import "net/http"

type Config struct {
	Host string
	Port string
}

type Server struct {
	Config  *Config
	Handler http.Handler
	routes  map[string]http.Handler
}

func NewServer(config *Config) *Server {
	return &Server{
		Config:  config,
		Handler: http.DefaultServeMux,
		routes:  make(map[string]http.Handler),
	}
}

func (s *Server) Run(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>База данных</h1"))
	})
}
