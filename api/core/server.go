package core

import (
	"net/http"
)

var Srv = &http.Server{}

func InitHttpServer(listeningAddress string, mux *http.ServeMux) {
	Srv.Addr = listeningAddress
	Srv.Handler = mux
}
