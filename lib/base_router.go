package lib

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type RouterHandle struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (r RouterHandle) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return r.Router.HandleFunc(path, handler).Methods("POST")
}

func (r RouterHandle) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return r.Router.HandleFunc(path, handler).Methods("GET")
}

func (r RouterHandle) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return r.Router.HandleFunc(path, handler).Methods("PUT")
}

func (r RouterHandle) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return r.Router.HandleFunc(path, handler).Methods("DELETE")
}

func (r RouterHandle) OPTIONS(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return r.Router.HandleFunc(path, handler).Methods("DELETE")
}

type RouterService interface {
	POST(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route
	GET(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route
	PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route
	DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route
	OPTIONS(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route
}

func NewRouterService(db *gorm.DB, r *mux.Router) RouterHandle {
	return RouterHandle{
		Router: r,
		DB:     db,
	}
}
