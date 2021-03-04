package lib

import "net/http"

type BaseHandler interface {
	Store(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	//Router(r *web.RouterHandle)
}

// swagger:parameters ParameterPage
type Parameter struct {
	Page int `json:"page"`
}
