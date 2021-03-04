package infrastructure

import (
	"github.com/gorilla/mux"
	"github.com/pujexx/go-boilerplate/lib"
	"net/http"
)

func NewPingHandler(route *mux.Router) {
	route.HandleFunc("/products", Ping).Methods("POST")
}
func Ping(w http.ResponseWriter, r *http.Request) {
	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: "Ping Pong",
	}, w, r)
	return
}
