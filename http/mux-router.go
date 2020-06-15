package http

import (
	"net/http"
	"fmt"
	
	"github.com/gorilla/mux"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")
}

func (*muxRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}