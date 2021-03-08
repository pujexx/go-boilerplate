package infrastructure

import (
	"github.com/gorilla/mux"
	"github.com/pujexx/go-boilerplate/lib"
	"net/http"
)

func NewPingHandler(route *mux.Router) {
	route.HandleFunc("/products", Ping).Methods("POST")
}
// ShowAccount godoc
// @Tag ping
// @Summary ping heath
// @Description ping heath
// @Accept  json
// @Success 200 {object} lib.Response
// @Header 200 {string} Token "qwerty"
// @Router /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: "Ping Pong",
	}, w, r)
	return
}
