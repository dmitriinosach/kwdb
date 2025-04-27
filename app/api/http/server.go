package http

import (
	"kwdb/app"
	"net/http"
	"strconv"
)

type cnf struct {
	host string
	port int
	soc  string
}

type Srv struct {
	config  *cnf
	handler http.Handler
	server  *http.Server
	routes  map[string]http.Handler
}

func NewServer() *Srv {
	s := &Srv{
		config: &cnf{
			host: app.Config.Get("HttpHost").(string),
			port: app.Config.Get("HttpPort").(int),
			soc:  app.Config.Get("HttpHost").(string) + ":" + strconv.Itoa(app.Config.Get("HttpPort").(int)),
		},
		handler: http.DefaultServeMux,
		routes:  make(map[string]http.Handler),
	}

	return s
}

func (s *Srv) Run(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>База данных</h1"))
	})
}
