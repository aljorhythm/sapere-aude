package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {

}

func NewCollectionServer(r mux.Router) {
	r.HandleFunc("/", rootHandler)
}
